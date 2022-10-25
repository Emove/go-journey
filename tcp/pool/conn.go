package pool

type Connection interface {
	Value(k interface{}) interface{}
	WithValue(k, v interface{})
}
