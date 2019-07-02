// A simple test helper
// see: https://github.com/google/googletest/blob/master/googletest/docs/primer.md
package assert

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

const (
	defaultSkip = 0
)

func ExpectEQ(t *testing.T, expected, actual interface{}, msg ...string) {
	if ret := equal(defaultSkip+1, expected, actual, msg...); ret != "" {
		fmt.Println(ret)
		t.Fail()
	}
}

func ExpectErr(t *testing.T, err error, msg ...string) {
	if ret := test(defaultSkip+1,
		fmt.Sprintf("Expected: error, actual:  '%v'", err),
		func() bool {
			return err != nil
		},
		msg...); ret != "" {
		fmt.Println(ret)
		t.Fail()
	}
}

func ExpectNoErr(t *testing.T, err error, msg ...string) {
	if ret := test(defaultSkip+1,
		fmt.Sprintf("Expected: no error, actual:  '%v'", err),
		func() bool {
			return err == nil
		},
		msg...); ret != "" {
		fmt.Println(ret)
		t.Fail()
	}
}

func ExpectTrue(t *testing.T, condition bool, msg ...string) {
	if ret := equal(defaultSkip+1, true, condition, msg...); ret != "" {
		fmt.Println(ret)
		t.Fail()
	}
}

func ExpectFalse(t *testing.T, condition bool, msg ...string) {
	if ret := equal(defaultSkip+1, false, condition, msg...); ret != "" {
		fmt.Println(ret)
		t.Fail()
	}
}

func ExpectNE(t *testing.T, expected, actual interface{}, msg ...string) {
	if ret := notEqual(defaultSkip+1, expected, actual, msg...); ret != "" {
		fmt.Println(ret)
		t.Fail()
	}
}

// Less Than
func ExpectLT(t *testing.T, val1, val2 interface{}, msg ...string) {
	if ret := compare(defaultSkip+1, val1, val2, "<", func() bool {
		if val1 == nil || val2 == nil {
			return false
		}
		//
		//s1, ok1 := val1.(sort.Interface)
		//s2, ok2 := val2.(sort.Interface)
		//if ok1 && ok2 {
		//	return s1.Less()
		//}
		// prime type
		v1 := reflect.ValueOf(val1)
		v2 := reflect.ValueOf(val2)
		if v1.Type() != v2.Type() {
			return false
		}
		switch v := val1.(type) {
		case int:
			return v < val2.(int)
		case int8:
			return v < val2.(int8)
		case int32:
			return v < val2.(int32)
		case int64:
			return v < val2.(int64)
		case uint:
			return v < val2.(uint)
		case uint8:
			return v < val2.(uint8)
		case uint32:
			return v < val2.(uint32)
		case uint64:
			return v < val2.(uint64)
		case float32:
			return v < val2.(float32)
		case float64:
			return v < val2.(float64)
		case string:
			return v < val2.(string)
		default:
			// not expected
			return false
		}
	}, msg...); ret != "" {
		fmt.Println(ret)
		t.Fail()
	}
}

//TODO(yu):
// ExpectLE
// ExpectGT
// ExpectGE

func compare(skip int, val1, val2 interface{}, operator string, cmp func() bool, msg ...string) string {
	return test(skip+1,
		fmt.Sprintf("Expected: '%v' %v '%v', actual:  '%v' (%v) vs. '%v' (%v)",
			val1, operator, val2,
			val1, reflect.TypeOf(val1), val2, reflect.TypeOf(val2)),
		cmp,
		msg...)
}

func equal(skip int, expected, actual interface{}, msg ...string) string {
	return test(skip+1,
		fmt.Sprintf("\tExpected: '%v'\n\tWhich is: %v\nTo be equal to: '%v'\n\tWhich is: %v",
			expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual)),
		func() bool {
			return reflect.DeepEqual(expected, actual)
		},
		msg...)
}

func notEqual(skip int, expected, actual interface{}, msg ...string) string {
	return test(skip+1,
		fmt.Sprintf("Expected: '%v' != '%v', actual:  '%v' (%v) vs. '%v' (%v)",
			expected, actual,
			expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual)),
		func() bool {
			return !reflect.DeepEqual(expected, actual)
		},
		msg...)
}

func test(skip int, concl string, cmp func() bool, msg ...string) string {
	if !cmp() {
		return fail(skip+1, "Failure:\n%s\nMessage: %s", concl, strings.Join(msg, " "))
	}
	return ""
}

func fail(skip int, format string, args ...interface{}) string {
	_, file, line, _ := runtime.Caller(skip + 1)
	return fmt.Sprintf("    %s:%d: %s\n", filepath.Base(file), line, fmt.Sprintf(format, args...))
}
