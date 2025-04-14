package logic

import (
	"errors"
	"fmt"
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
	}  else if succedent[implicationOperator+1] == "(" {
		// fmt.Println(succedent)
		indexOfFirstBracket := slices.Index(succedent, "(")
		slices.Reverse(succedent)
		// fmt.Println(succedent)
		indexOfLastBracket := len(succedent) - slices.Index(succedent, ")") - 1
		// fmt.Printf("index of last ): %v\n", len(succedent)-index)
		slices.Reverse(succedent)
		// fmt.Println(succedent[indexOfLastBracket])
		fmt.Printf("index 1: %v, index 2: %v\n", indexOfFirstBracket, indexOfLastBracket)
		// fmt.Println(indexOfLastBracket)
		bracketedContent := succedent[indexOfFirstBracket+1 : indexOfLastBracket]

		leftMost = append([]string{}, bracketedContent...)
		// fmt.Printf("leftMost: %v\n", leftMost)
	} else {
		leftMost = append([]string{}, succedent[implicationOperator+1])
	}

	if succedent[implicationOperator+1] == Negation {
		rightMost = succedent[implicationOperator+3:]
	} else if succedent[implicationOperator+1] == "(" {
		slices.Reverse(succedent)
		indexOfLastBracket := len(succedent) - slices.Index(succedent, ")") - 1
		slices.Reverse(succedent)
		rightMost = succedent[indexOfLastBracket+1:]
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

	var leftMost []string
	var rightMost []string

	if !slices.Contains(antecedent, "→") {
		return [][]string{}, [][]string{}, errors.New("no implication")
	}

	implicationOperator := slices.Index(antecedent, "→")

	if implicationOperator == -1 {
		return [][]string{}, [][]string{}, errors.New("index out of range")
	}

	if (antecedent[implicationOperator+1] == "(") {
		indexOfFirstBracket := slices.Index(antecedent, "(")
		slices.Reverse(antecedent)
		indexOfLastBracket := len(antecedent) - slices.Index(antecedent, ")") - 1
		slices.Reverse(antecedent)
		// fmt.Printf("index 1: %v, index 2: %v\n", indexOfFirstBracket, indexOfLastBracket)
		bracketedContent := antecedent[indexOfFirstBracket +1 : indexOfLastBracket]
		leftMost = bracketedContent

		rightMost = antecedent[indexOfLastBracket+1:]
	}

	// leftMost := antecedent[implicationOperator+1]
	// rightMost := antecedent[implicationOperator+2]

	// build left
	leftSuccedent := append(leftMost, succedent...)
	leftAntecedent := append([]string{}, antecedent[:implicationOperator]...)
	leftSequent := append([][]string{}, leftAntecedent, leftSuccedent)

	fmt.Println(leftSequent);

	// build
	rightSuccedent := append([]string{}, succedent...)
	rightAntecedent1 := append([]string{}, antecedent[:implicationOperator]...)
	rightAntecedent2 := append(rightAntecedent1, rightMost...)
	rightSequent := append([][]string{}, rightAntecedent2, rightSuccedent)

	fmt.Println(rightSequent)

	return leftSequent, rightSequent, nil
}
