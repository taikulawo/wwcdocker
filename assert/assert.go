// assert 包提供基础的断言能力
// 后期考虑将整个包独立出来，加入我的通用工具包，避免重复开发
package assert

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

/*
* Golang 不支持泛型，用接口实现很难看
* 2.0草稿引入的 handle 与 check 只是一种 error 的处理方案。
* 我所理解是的让代码不那么难看，但并未引入一种可以打断当前 control flow 的能力
* 或者说， Go 只是让你处理 error 更容易，但却并未引入异常机制（可以直接打算当前 control flow，一直回溯到上层能够处理的位置）
* 而新的 Try 提案也因为社区的反对而搁浅。
* 听取社区的意见固然是好事，但在一些方面一定要果断。漫天的讨论没有意义
 */

type Matcher struct {
	method reflect.Value
	verb   string
}

func (m *Matcher) call(value interface{}, expect ...interface{}) bool {
	input := make([]reflect.Value, m.method.Type().NumIn())
	input[0] = reflect.ValueOf(value)
	for index, v := range expect {
		input[index+1] = reflect.ValueOf(v)
	}
	result := m.method.Call(input)
	return result[0].Bool()
}

var (
	equals = make(map[reflect.Type]*Matcher)
)

func createMatcher(fn interface{}, verb string) *Matcher {
	value := reflect.ValueOf(fn)
	if value.Kind() != reflect.Func {
		panic("fn is not a function")
	}
	return &Matcher{
		method: value,
	}
}

func storeToTable(m map[reflect.Type]*Matcher, fn reflect.Value, matcher *Matcher) {
	t := fn.Type().In(0)
	m[t] = matcher
}

func registerEqualsMatcher(fn interface{}, verb string) {
	matcher := createMatcher(fn, verb)
	// todo reflect.ValueOf(fn)
	storeToTable(equals, reflect.ValueOf(fn), matcher)
}

func IsEquals(a, b interface{}) bool {
	t := reflect.TypeOf(b)
	return equals[t].call([]interface{}{a, b})
}

func init() {
	registerEqualsMatcher(func(a, b int) bool {
		return a == b
	}, "equals to")
	registerEqualsMatcher(func(a, b string) bool {
		return a == b
	}, "equals to")
}

type Assert func(value interface{}, matcher *Matcher, expect ...interface{})

func With(t *testing.T) Assert {
	return func(value interface{}, matcher *Matcher, expect ...interface{}) {
		if !matcher.call(value, expect) {
			// 打印错误信息
			msg := []interface{}{"Not true that (", value, ")", matcher.verb, "( "}
			msg = append(msg, expect...)
			msg = append(msg, ")")
			fmt.Println(msg)
			t.FailNow()
		}
	}
}

// getCaller return file name and line number
func getCaller() (string, int) {
	stackLevel := 1
	for {
		_, file, line, ok := runtime.Caller(stackLevel)
		if strings.Contains(file, "assert") {
			stackLevel++
		} else {
			if ok {
				if index := strings.LastIndex(file, "/"); index >= 0 {
					file = file[index+1:]
				} else if index := strings.LastIndex(file, "\\"); index >= 0 {
					file = file[index+2:]
				}
			} else {
				file = "??"
				line = 1
			}
			return file, line
		}
	}
}

// decorate 格式化输出。
func decorate(s string) string {
	file, line := getCaller()
	buf := new(bytes.Buffer)
	// Every line is indented at least one tab.
	buf.WriteString("  ")
	fmt.Fprintf(buf, "%s:%d: ", file, line)
	lines := strings.Split(s, "\n")
	if l := len(lines); l > 1 && lines[l-1] == "" {
		lines = lines[:l-1]
	}
	for i, line := range lines {
		if i > 0 {
			buf.WriteString("\n\t\t")
		}
		buf.WriteString(line)
	}
	buf.WriteByte('\n')
	return buf.String()
}
