package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"strconv"
	"strings"
	"unicode"
)

func parse() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "D:\\workspace\\profile\\go\\go-journey\\ast\\model\\user.go", nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if stru, ok := x.Type.(*ast.StructType); ok {
				model := parseModel(x.Name.Name, stru)
				jsonBytes, _ := json.Marshal(model)
				fmt.Println(string(jsonBytes))
				for _, field := range model.Fields {
					if field.Name != "Gender" {
						continue
					}
					field.BuildParserSegment(model.Name)
				}
			}
		}
		return true
	})
}

const (
	TagKeyGorm   = "gorm"
	TagKeyParser = "parser"
)

type Tag struct {
	Values map[string]string `json:"values,omitempty"`
}

func NewTag() *Tag {
	return &Tag{make(map[string]string)}
}

func (t *Tag) AddPair(key, value string) {
	t.Values[key] = value
}

func (t *Tag) Put(key string) {
	t.Values[key] = ""
}

type Model struct {
	Name        string        `json:"name"`
	InheritForm []string      `json:"inherit_form"`
	Fields      []*ModelField `json:"fields"`
}

type ModelField struct {
	Name       string `json:"name"`
	GormTags   *Tag   `json:"gorm_tags,omitempty"`
	ParserTags *Tag   `json:"parser_tags,omitempty"`
	Type       string `json:"type"`
}

func (mf *ModelField) BuildParserSegment(modelName string) {
	if mf.ParserTags == nil || len(mf.GormTags.Values) == 0 {
		return
	}
	tl := `
type {{.ModelName}}Repo interface {
	FindByID(ctx context.Context, {{.PrefixName}} *{{.ModelName}}) error
	FindByIDs(ctx context.Context, ids []int) ([]*{{.ModelName}}, error)
	Select(ctx context.Context, {{.PrefixName}} *{{.ModelName}}, page *data.Page, orders []*data.Order) ([]*{{.ModelName}}, *data.Page, error)
	Create(ctx context.Context, {{.PrefixName}} *{{.ModelName}}) (bool, error)
	Save(ctx context.Context, {{.PrefixName}} *{{.ModelName}}) (bool, error)
	UpdateByIDAndChanges(ctx context.Context, {{.PrefixName}} *{{.ModelName}}) (bool, error)
	Delete(ctx context.Context, id int) error
}

type {{.Name}} int
const (
	{{range $key, $val := .EnumValue}}
		{{$key}} = {{$val}}
	{{end}}
	{{range $key, $val := .EnumName}}
		{{$key}} = "{{$val}}"
	{{end}}
)
func (t {{.Name}}) String() string {
	values := map[int]struct{}{
		{{range $idx, $val := .EnumLabels}}
			int({{$val}}): {},
		{{end}}
	}
	if _, ok := values[int(t)]; !ok {
		return fmt.Sprintf("WorkstationRuntimeTaskType(%d)", t)
	}
	types := [...]string{
		{{range $key, $val := .EnumNames}}
			{{$val}}
		{{end}}
	}
	return types[t-1]
}

type {{.ModelName}}Usecase struct {
	tm     data.Transaction
	repo   {{.ModelName}}Repo
	logger *log.Helper
}

func New{{.ModelName}}Usecase(tm data.Transaction, repo {{.ModelName}}Repo, logger log.Logger) *{{.ModelName}}Usecase {
	return &{{.ModelName}}Usecase{
		tm:     tm,
		repo:   repo,
		logger: log.NewHelper(logger),
	}
}

{{range $key, $val := .Fields}}
func (m *{{$val.ModelName}}) Set{{$val.Name}}(value {{$val.Type}}) {
	m.{{$val.Name}} = value
	m.Update("{{$val.Column}}", value)
}
{{end}}

`
	tpl, err := template.New("").Parse(tl)
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"ModelName":  modelName,
		"PrefixName": LowerCamelCae(modelName),
		"Name":       fmt.Sprintf("%s%s", modelName, UpperCamelCase(mf.Name)),
	}
	for k, v := range mf.ParserTags.Values {
		if k == "enum" {
			mapping := parseEnumValue(v)
			enumValue := map[string]interface{}{}
			enumName := map[string]string{}
			var enumNames []string
			var enumLabels []string
			for label, value := range mapping {
				enumLabel := fmt.Sprintf("%s%s%s", UpperCamelCase(modelName), UpperCamelCase(mf.Name), UpperCamelCase(label))
				enumLabels = append(enumLabels, enumLabel)
				enumValue[enumLabel] = value
				labelName := fmt.Sprintf("%s%s%sName", UpperCamelCase(modelName), UpperCamelCase(mf.Name), UpperCamelCase(label))
				enumName[labelName] = UpperCamelCase(label)
				enumNames = append(enumNames, labelName)
			}
			data["EnumValue"] = enumValue
			data["EnumName"] = enumName
			data["EnumLabels"] = enumLabels
			data["EnumNames"] = enumNames
		}
	}
	type InternalFiled struct {
		ModelName string
		Name      string
		Column    string
		Type      string
	}
	var columnName string
	for k, v := range mf.GormTags.Values {
		if k == "column" {
			columnName = v
		}
	}
	fileds := []InternalFiled{
		{
			ModelName: modelName,
			Name:      mf.Name,
			Column:    columnName,
			Type:      mf.Type,
		},
	}
	data["Fields"] = fileds
	generateBuffer := new(bytes.Buffer)
	if err = tpl.Execute(generateBuffer, data); err != nil {
		panic(err)
	}
	fmt.Println(string(generateBuffer.Bytes()))
}

func parseEnumValue(v string) map[string]string {
	vs := strings.Split(strings.TrimSpace(v), ",")
	result := make(map[string]string)
	for _, enumPair := range vs {
		pair := strings.Split(enumPair, "-")
		result[pair[1]] = pair[0]
	}
	return result
}

func parseModel(modelName string, stru *ast.StructType) *Model {
	if !isInheritFromModel(stru) {
		return nil
	}
	fields := make([]*ModelField, 0, len(stru.Fields.List)-1)

	for _, field := range stru.Fields.List {
		switch field.Type.(type) { // 暂时不处理指针类型
		case *ast.Ident:
			fields = append(fields, &ModelField{
				Name:       field.Names[0].Name,
				GormTags:   parseTagValue(field.Tag.Value, TagKeyGorm),
				ParserTags: parseTagValue(field.Tag.Value, TagKeyParser),
				Type:       getFieldTypeName(field),
			})
		}

	}
	//jsonStr, _ := json.Marshal(fields)
	//fmt.Println(string(jsonStr))
	return &Model{Name: modelName, InheritForm: []string{"Model"}, Fields: fields}
}

func getFieldTypeName(field *ast.Field) string {
	if field.Type == nil {
		return ""
	}
	switch x := field.Type.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.StarExpr:
		return x.X.(*ast.Ident).Name
	}
	return ""
}

func isInheritFromModel(stru *ast.StructType) bool {
	if stru.Fields == nil || len(stru.Fields.List) == 0 {
		return false
	}
	if v, ok := (stru.Fields.List[0].Type).(*ast.StarExpr); ok {
		m := v.X.(*ast.Ident)
		if m.Name == "Model" {
			fmt.Println("struct is inherit from model")
			return true
		}
	}
	return false
}

func parseTagValue(value, expectedScope string) *Tag {
	scopes := strings.Split(strings.Trim(value, "`"), " ")
	result := NewTag()
	for _, scope := range scopes {
		if len(scope) <= len(expectedScope) {
			continue
		}
		if scope[:len(expectedScope)] != expectedScope {
			continue
		}
		tags := strings.Split(scope[len(expectedScope)+1:], ";")
		for _, tag := range tags {
			tagVal, _ := strconv.Unquote(tag)
			if strings.Contains(tagVal, ":") {
				pair := strings.Split(tagVal, ":")
				result.AddPair(pair[0], pair[1])
			} else {
				result.Put(tagVal)
			}
		}
	}
	return result
}

func UpperCamelCase(str string) string {
	if unicode.IsUpper(rune(str[0])) {
		return str
	}
	return strings.ToUpper(str[:1]) + str[1:]
}

func LowerCamelCae(str string) string {
	if unicode.IsLower(rune(str[0])) {
		return str
	}
	return strings.ToLower(str[:1]) + str[1:]
}

func main() {
	parse()
}
