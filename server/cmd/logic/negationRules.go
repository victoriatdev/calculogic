package logic

import (
	"fmt"
	"fyp-server/cmd/utils"
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

	fmt.Println(negationOperator)

	if negationOperator == 0 {
		if succedent[negationOperator+1] == "(" {
			firstIndex, lastIndex := utils.HandleBrackets(succedent)
			leftMost := succedent[firstIndex:lastIndex]
			delta := succedent[lastIndex+1:]
			antecedent := append(leftMost, gamma...)
			sequent := append([][]string{}, antecedent, delta)

			return sequent
		} else {
			leftMost := succedent[negationOperator+1]
			delta := succedent[negationOperator+2:]
			antecedent := append([]string{leftMost}, gamma...)
			sequent = append([][]string{}, antecedent, delta)
		}
	} else {
		// fmt.Println(sequent)
		pre := succedent[:negationOperator]
		// fmt.Println(pre)
		if succedent[negationOperator+1] == "(" {
			firstIndex, lastIndex := utils.HandleBrackets(succedent)
			leftMost := succedent[firstIndex:lastIndex]
			delta := succedent[lastIndex+1:]
			delta = append(pre, delta...)
			antecedent := append(leftMost, gamma...)
			sequent = append([][]string{}, antecedent, delta)
		} else {
			leftMost := succedent[negationOperator+1]
			// fmt.Println(leftMost)
			delta := succedent[negationOperator+2:]
			delta = append(pre, delta...)
			// fmt.Println(delta)
			antecedent := append([]string{leftMost}, gamma...)
			sequent = append([][]string{}, antecedent, delta)
		}
	}

	return sequent
}

func ApplyNegationLeft(sequent [][]string) [][]string {

	var antecedent = sequent[0]
	var delta = sequent[1]

	if !slices.Contains(antecedent, "¬") {
		return sequent
	}

	negationOperator := slices.Index(antecedent, "¬")

	if antecedent[negationOperator+1] == "(" {
		firstIndex, lastIndex := utils.HandleBrackets(antecedent)
		rightMost := antecedent[firstIndex:lastIndex]
		gamma := antecedent[:negationOperator]
		succedent := append(rightMost, delta...)
		sequent = append([][]string{}, gamma, succedent)
		return sequent
	}

	rightMost := antecedent[negationOperator+1]
	gamma := antecedent[:negationOperator]

	succedent := append([]string{rightMost}, delta...)

	sequent = append([][]string{}, gamma, succedent)

	return sequent
}
