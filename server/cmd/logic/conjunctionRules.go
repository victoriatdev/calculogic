package logic

import (
	"errors"
	"fmt"
	"slices"
)

// CONJUNCTION RULES
func ApplyConjunctionRight(sequent [][]string) ([][]string, [][]string, error) {

	gamma := sequent[0]
	succedent := sequent[1]

	// fmt.Println(slices.Contains(succedent, Conjunction))
	// fmt.Println(succedent)

	if !slices.Contains(succedent, "∧") {
		fmt.Println("No conjnction")
		return sequent, [][]string{}, errors.New("no conjunction")
	}

	conjunctionOperator := slices.Index(succedent, "∧")
	if conjunctionOperator == -1 {
		return [][]string{}, [][]string{}, errors.New("index out of range")
	}
	// leftMost := succedent[conjunctionOperator+1]
	// rightMost := succedent[conjunctionOperator+2]
	// delta := succedent[conjunctionOperator+3:]

	var leftMost []string
	var rightMost []string
	var delta []string

	if succedent[conjunctionOperator+1] == Negation {
		leftMost = succedent[conjunctionOperator+1 : conjunctionOperator+3]
		fmt.Println(leftMost)
		rightMost = succedent[conjunctionOperator+3:]
		fmt.Println(rightMost)
		delta = succedent[conjunctionOperator+4:]
	} else {
		leftMost = append([]string{}, succedent[conjunctionOperator+1])
		rightMost = append([]string{}, succedent[conjunctionOperator+2])
		delta = succedent[conjunctionOperator+3:]
	}

	// build leftSequent

	leftSuccedent := append(append([]string{}, leftMost...), delta...)
	leftSequent := append([][]string{}, gamma, leftSuccedent)

	// build rightSequent

	rightSuccedent := append(append([]string{}, rightMost...), delta...)
	rightSequent := append([][]string{}, gamma, rightSuccedent)

	return leftSequent, rightSequent, nil
}

func ApplyConjunctionLeft(sequent [][]string) ([][]string, error) {

	antecedent := sequent[0]
	delta := sequent[1]

	if !slices.Contains(antecedent, "∧") {
		return sequent, errors.New("no conjunction")
	}

	conjunctionOperator := slices.Index(antecedent, "∧")
	if conjunctionOperator == -1 {
		return [][]string{}, errors.New("index out of range")
	}

	tmpAntecedent := append([]string{}, antecedent[:conjunctionOperator]...)
	tmpAntecedent = append(tmpAntecedent, antecedent[conjunctionOperator+1:]...)

	sequent = append([][]string{}, tmpAntecedent, delta)

	return sequent, nil
}
