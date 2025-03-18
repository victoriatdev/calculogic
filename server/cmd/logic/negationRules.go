package logic

import (
	"slices"
)

// NEGATION RULES
func ApplyNegationRight(sequent [][]string) [][]string {

	var gamma = sequent[0]
	var succedent = sequent[1]

	if !slices.Contains(succedent, "¬") {
		return sequent
	}

	negationOperator := slices.Index(succedent, "¬")
	leftMost := succedent[negationOperator+1]
	delta := succedent[negationOperator+2:]

	antecedent := append([]string{leftMost}, gamma...)

	sequent = append([][]string{}, antecedent, delta)

	return sequent
}

func ApplyNegationLeft(sequent [][]string) [][]string {

	var antecedent = sequent[0]
	var delta = sequent[1]

	if !slices.Contains(antecedent, "¬") {
		return sequent
	}

	negationOperator := slices.Index(antecedent, "¬")
	rightMost := antecedent[negationOperator+1]
	gamma := antecedent[:negationOperator]

	succedent := append([]string{rightMost}, delta...)

	sequent = append([][]string{}, gamma, succedent)

	return sequent
}
