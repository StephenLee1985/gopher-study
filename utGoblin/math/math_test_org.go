package math

import (
	"testing"
)

func Test_AddInt(t *testing.T) {

	sum := AddInt(1, 3)
	if sum != 4 {
		t.Error("1+3 should be 4")
	}
}
