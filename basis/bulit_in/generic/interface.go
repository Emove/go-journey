package generic

import (
	"strconv"
)

type Type[T any] interface {
	Type() string
	IsValid(string) bool
	Stringify(T) string
	Convert(string) (T, error)
}

type Builtin interface {
}

type boolType[T bool] struct {
}

func (*boolType[T]) Type() string {
	return "Bool"
}

func (*boolType[T]) IsValid(value string) bool {
	ok, err := strconv.ParseBool(value)
	if err != nil || !ok {
		return false
	}
	return true
}

func (*boolType[T]) Stringify(value bool) string {
	return strconv.FormatBool(value)
}

func (*boolType[T]) Convert(value string) (bool, error) {
	return strconv.ParseBool(value)
}

type intType[T int] struct {
}

func (*intType[T]) Type() string {
	return "Int"
}

func (*intType[T]) IsValid(value string) bool {
	_, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return false
	}
	return true
}

func (*intType[T]) Stringify(value int) string {
	return strconv.FormatInt(int64(value), 10)
}

func (*intType[T]) Convert(value string) (int, error) {
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(result), nil
}

var types [][]Type[any]

func IsValidValue(valueType, value string) bool {
	types = append(types, []Type[bool]{&boolType[bool]{}})
	for _, validator := range types {
		if validator.Type() == valueType {
			return validator.IsValid(value)
		}
	}
	return false
}
