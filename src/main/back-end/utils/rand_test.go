package utils

import "testing"

/**
    utils
    @author: roccoshi
    @desc: Test rand.go
**/

func TestIRand_Int(t *testing.T) {
	x := Rand.Int(1, 200000)
	y := Rand.Int(-1000000, 0)
	t.Logf("x = %v, y = %v", x, y)
	if x == y {
		t.Fail()
	}
}

func TestIRand_String(t *testing.T) {
	x := Rand.String(32)
	y := Rand.String(40)
	t.Logf("x = %v, y = %v", x, y)
	if len(x) != 32 || len(y) != 40 || x == y {
		t.Fail()
	}
}
