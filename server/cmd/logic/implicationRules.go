package logic

import (
	"errors"
	"slices"
)

// IMPLICATION RULES
func ApplyImplicationRight(sequent [][]string) ([][]string, error) {

	var antecedent = sequent[0]
	var succedent = sequent[1]

	// determine succedent values
	// [[], [->, P, P]]

	if !slices.Contains(succedent, "→") {
		return [][]string{}, errors.New("no implication")
	}

	implicationOperator := slices.Index(succedent, "→")

	if implicationOperator == -1 {
		return [][]string{}, errors.New("index out of range")
	}

	var leftMost []string
	var rightMost []string

	if succedent[implicationOperator+1] == Negation {
		leftMost = succedent[implicationOperator+1 : implicationOperator+2]
	} else {
		leftMost = append([]string{}, succedent[implicationOperator+1])
	}

	if succedent[implicationOperator+1] == Negation {
		rightMost = succedent[implicationOperator+3:]
	} else {
		rightMost = succedent[implicationOperator+2:]
	}

	// rightMost := succedent[implicationOperator+2:]

	// prepend leftmost to antecedent via composite literal

	antecedent = append(append([]string{}, leftMost...), antecedent...)

	sequent = append([][]string{}, antecedent, rightMost)

	return sequent, nil
}

func ApplyImplicationLeft(sequent [][]string) ([][]string, [][]string, error) {

	antecedent := sequent[0]
	succedent := sequent[1]

	if !slices.Contains(antecedent, "→") {
		return [][]string{}, [][]string{}, errors.New("no implication")
	}

	implicationOperator := slices.Index(antecedent, "→")

	if implicationOperator == -1 {
		return [][]string{}, [][]string{}, errors.New("index out of range")
	}

	leftMost := antecedent[implicationOperator+1]
	rightMost := antecedent[implicationOperator+2]

	// build left
	leftSuccedent := append([]string{leftMost}, succedent...)
	leftAntecedent := append([]string{}, antecedent[:implicationOperator]...)
	leftSequent := append([][]string{}, leftAntecedent, leftSuccedent)

	// build
	rightSuccedent := append([]string{}, succedent...)
	rightAntecedent := append(antecedent[:implicationOperator], rightMost)
	rightSequent := append([][]string{}, rightAntecedent, rightSuccedent)

	return leftSequent, rightSequent, nil
}
