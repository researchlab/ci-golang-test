package cigolangtest

import (
	"testing"
)

func TestReturnStr(t *testing.T) {
	if ReturnStr() != "ABC" {
		t.Errorf("it expected 'ABC' but %v", ReturnStr())
	}
}
