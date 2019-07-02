package assert

import (
	"errors"
	"testing"
)

func TestExpectEQ(t *testing.T) {
	ExpectEQ(t, "aaa", "aaa", "message for test equal")
}

func TestExpectNE(t *testing.T) {
	ExpectNE(t, "aaa", 1, "message for test not equal")
}

func TestExpectTrue(t *testing.T) {
	ExpectTrue(t, true)
}

func TestExpectFalse(t *testing.T) {
	ExpectFalse(t, false)
}

func TestExpectLT(t *testing.T) {
	ExpectLT(t, "aaa", "aab")
}

func TestExpectErr(t *testing.T) {
	ExpectErr(t, errors.New("fake error"))
}

func TestExpectNoErr(t *testing.T) {
	ExpectNoErr(t, nil)
}
