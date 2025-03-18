package logic

import (
	"slices"
)

// DISJUNCTION RULES
func ApplyDisjunctionLeft(sequent [][]string) ([][]string, [][]string) {

	antecedent := sequent[0]
	delta := sequent[1]

	if !slices.Contains(antecedent, "∨") {
		return sequent, [][]string{}
	}

	disjunctionOperator := slices.Index(antecedent, "∨")
	leftMost := antecedent[disjunctionOperator+1]
	rightMost := antecedent[disjunctionOperator+2]
	gamma := antecedent[:disjunctionOperator]

	// build leftSequent
	leftAntecedent := append(append([]string{}, gamma...), leftMost)
	leftSequent := append([][]string{}, leftAntecedent, delta)

	// build rightSequent

	rightAntecedent := append(append([]string{}, gamma...), rightMost)
	rightSequent := append([][]string{}, rightAntecedent, delta)

	return leftSequent, rightSequent
}

func ApplyDisjunctionRight(sequent [][]string) [][]string {
	gamma := sequent[0]
	succedent := sequent[1]

	if !slices.Contains(succedent, "∨") {
		return sequent
	}

	conjunctionOperator := slices.Index(succedent, "∨")

	tmpSuccedent := append([]string{}, succedent[:conjunctionOperator]...)

	tmpSuccedent = append(tmpSuccedent, succedent[conjunctionOperator+1:]...)

	sequent = append([][]string{}, gamma, tmpSuccedent)

	return sequent
}
