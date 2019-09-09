package assert

import "reflect"

type Matcher struct {
	method reflect.Value
}

func (m *Matcher) call(args []interface{}) bool{
	input := make([]reflect.Value, m.method.Type().NumIn())
	for index, v := range args {
		input[index] = reflect.ValueOf(v)
	}
	result := m.method.Call(input)
	return result[0].Bool()
}

var (
	equals = make(map[reflect.Type]*Matcher)
)

func createMatcher(fn interface{})*Matcher {
	value := reflect.ValueOf(fn)
	if value.Kind() != reflect.Func {
		panic("fn is not a function")
	}
	return & Matcher {
		method: value,
	}
}

func storeToTable(m map[reflect.Type]*Matcher, fn reflect.Value, matcher *Matcher) {
	// t := reflect.TypeOf(fn).In(0)
	t := fn.Type().In(0)
	m[t] = matcher
}

func registerEqualsMatch(fn interface{}) {
	matcher := createMatcher(fn)
	// todo reflect.ValueOf(fn)
	storeToTable(equals, reflect.ValueOf(fn), matcher)
}

func IsEquals(a,b interface{}) bool{
	t := reflect.TypeOf(b)
	return equals[t].call([]interface{}{a,b})
}


func init() {
	registerEqualsMatch(func(a,b int) bool{
		return a == b
	})
}