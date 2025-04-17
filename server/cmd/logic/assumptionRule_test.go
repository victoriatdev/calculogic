package logic

import (
	"testing"
)

func TestValidAssumption(t *testing.T) {
	validAssumption := [][]string{{"A"},{"A"}}
	want := true
	res := ApplyAssumption(validAssumption)
	if res != want {
		t.Errorf("ApplyAssumption(%v) = %v, wanted %v", validAssumption, res, want)
	}
}

func TestInalidAssumption(t *testing.T) {
	validAssumption := [][]string{{"A"},{"B"}}
	want := false
	res := ApplyAssumption(validAssumption)
	if res != want {
		t.Errorf("ApplyAssumption(%v) = %v, wanted %v", validAssumption, res, want)
	}
}