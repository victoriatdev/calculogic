package logic

import (
	"slices"
)

func ApplyAssumption(sequent [][]string) bool {

	// check left and right arrays for a common singular element without a logical operator
	var antecedent = sequent[0]
	var succedent = sequent[1]

	if slices.Contains(antecedent, "→") || slices.Contains(antecedent, "¬") ||
		slices.Contains(antecedent, "∧") || slices.Contains(antecedent, "∨") {
		return false
	}

	if slices.Contains(succedent, "→") || slices.Contains(succedent, "¬") ||
		slices.Contains(succedent, "∧") || slices.Contains(succedent, "∨") {
		return false
	}

	for i := range len(antecedent) { //[Q,P]
		if slices.Contains(succedent, antecedent[i]) {
			return true
		}
	}

	return false
}




