package logic

import (
	"errors"
	"fmt"
	"slices"
)

// IMPLICATION RULES
func ApplyImplicationRight(sequent [][]string) ([][]string, error) {

	newSequent := make([][]string, len(sequent))
	copy(newSequent, sequent)

	var newAntecedent = newSequent[0]
	var newSuccedent = newSequent[1]

	// determine succedent values
	// [[], [->, P, P]]

	if !slices.Contains(newSuccedent, "→") {
		return [][]string{}, errors.New("no implication")
	}

	implicationOperator := slices.Index(newSuccedent, "→")

	if implicationOperator == -1 {
		return [][]string{}, errors.New("index out of range")
	}

	var leftMost []string
	var rightMost []string

	if newSuccedent[implicationOperator+1] == Negation {
		if newSuccedent[implicationOperator+2] == "(" {
			slices.Reverse(newSuccedent)
			indexOfLastBracket := len(newSuccedent) - slices.Index(newSuccedent, ")") - 1
			slices.Reverse(newSuccedent)
			bracketedContent := newSuccedent[implicationOperator+1 : indexOfLastBracket+1]
			leftMost = append([]string{}, bracketedContent...)
		}
	} else if newSuccedent[implicationOperator+1] == "(" {
		indexOfFirstBracket := slices.Index(newSuccedent, "(")
		slices.Reverse(newSuccedent)
		indexOfLastBracket := len(newSuccedent) - slices.Index(newSuccedent, ")") - 1
		slices.Reverse(newSuccedent)
		fmt.Printf("index 1: %v, index 2: %v\n", indexOfFirstBracket, indexOfLastBracket)
		bracketedContent := newSuccedent[indexOfFirstBracket+1 : indexOfLastBracket]

		leftMost = append([]string{}, bracketedContent...)
	} else {
		leftMost = append([]string{}, newSuccedent[implicationOperator+1])
	}

	if newSuccedent[implicationOperator+1] == Negation {
		if newSuccedent[implicationOperator+2] == "(" {
			slices.Reverse(newSuccedent)
			indexOfLastBracket := len(newSuccedent) - slices.Index(newSuccedent, ")") - 1
			slices.Reverse(newSuccedent)
			rightMost = newSuccedent[indexOfLastBracket+1:]
		} else {
			rightMost = newSuccedent[implicationOperator+3:]
		}
	} else if newSuccedent[implicationOperator+1] == "(" {
		slices.Reverse(newSuccedent)
		indexOfLastBracket := len(newSuccedent) - slices.Index(newSuccedent, ")") - 1
		slices.Reverse(newSuccedent)
		rightMost = newSuccedent[indexOfLastBracket+1:]
	} else {
		rightMost = newSuccedent[implicationOperator+2:]
	}

	// rightMost := newSuccedent[implicationOperator+2:]

	// prepend leftmost to antecedent via composite literal

	newAntecedent = append(append([]string{}, leftMost...), newAntecedent...)

	newSequent = append([][]string{}, newAntecedent, rightMost)

	return newSequent, nil
}

func ApplyImplicationLeft(sequent [][]string) ([][]string, [][]string, error) {

	antecedent := sequent[0]
	succedent := sequent[1]

	var leftMost []string
	var rightMost []string

	if !slices.Contains(antecedent, "→") {
		return [][]string{}, [][]string{}, errors.New("no implication")
	}

	implicationOperator := slices.Index(antecedent, "→")

	if implicationOperator == -1 {
		return [][]string{}, [][]string{}, errors.New("index out of range")
	}

	if antecedent[implicationOperator+1] == "(" {
		indexOfFirstBracket := slices.Index(antecedent, "(")
		slices.Reverse(antecedent)
		indexOfLastBracket := len(antecedent) - slices.Index(antecedent, ")") - 1
		slices.Reverse(antecedent)
		// fmt.Printf("index 1: %v, index 2: %v\n", indexOfFirstBracket, indexOfLastBracket)
		bracketedContent := antecedent[indexOfFirstBracket+1 : indexOfLastBracket]
		leftMost = bracketedContent

		rightMost = antecedent[indexOfLastBracket+1:]
	}

	// leftMost := antecedent[implicationOperator+1]
	// rightMost := antecedent[implicationOperator+2]

	// build left
	leftSuccedent := append(leftMost, succedent...)
	leftAntecedent := append([]string{}, antecedent[:implicationOperator]...)
	leftSequent := append([][]string{}, leftAntecedent, leftSuccedent)

	fmt.Println(leftSequent)

	// build
	rightSuccedent := append([]string{}, succedent...)
	rightAntecedent1 := append([]string{}, antecedent[:implicationOperator]...)
	rightAntecedent2 := append(rightAntecedent1, rightMost...)
	rightSequent := append([][]string{}, rightAntecedent2, rightSuccedent)

	fmt.Println(rightSequent)

	return leftSequent, rightSequent, nil
}
