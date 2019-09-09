package assert

import (
	"reflect"
	"testing"
)

func TestIntEquals(t *testing.T) {
	var a int
	var b int
	a = 1
	b = 1
	result := IsEquals(a, b)
	t.Log(result)
	s1 := "a"
	s2 := "a"
	result = IsEquals(s1, s2)
	a = 1
}

func TestTypeIsEquals(t *testing.T) {
	a := func(b interface{}) {

	}
	b := func(c int) {

	}

	t1 := reflect.TypeOf(a)
	t2 := reflect.TypeOf(b)
	r := t1 == t2
	t.Log(r)
}
