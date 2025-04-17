package logic

import (
	"errors"
	"slices"
	"testing"
)

func TestImplicationRightRule(t *testing.T) {
	valid := [][]string{{}, {"→", "A", "A"}}
	res, err := ApplyImplicationRight(valid)
	want := [][]string{{"A"}, {"A"}}

	if !slices.Equal(res[0], want[0]) && !slices.Equal(res[1], want[1]) && err != nil {
		t.Errorf("ApplyAssumption(%v) = %v, wanted %v, got error: %v", valid, res, want, err)
	}
}

func TestInvalidImplicationRule(t *testing.T) {
	input := [][]string{{"A"}, {"B"}}
	want := errors.New("no implication")
	res, err := ApplyImplicationRight(input)
	if err == nil {
		t.Errorf("ApplyImplicationRight(%v) = %v, wanted %v", input, res, want)
	}
}
