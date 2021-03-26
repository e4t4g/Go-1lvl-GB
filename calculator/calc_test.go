package calculator

import (
	"testing"
)

func TestCalc(t *testing.T) {
	if Calc(4,"+",6) != 10 {
		t.Error(`Calc ("4+6=10") = false`)
	}
}