package utils

import (
	"slices"
)

func HandleBrackets(inputList []string) (firstIndex int, lastIndex int) {
		indexOfFirstBracket := slices.Index(inputList, "(") + 1
		slices.Reverse(inputList)
		indexOfLastBracket := len(inputList) - slices.Index(inputList, ")") - 1
		slices.Reverse(inputList)
		// bracketedContent := inputList[indexOfFirstBracket+1 : indexOfLastBracket]

		return indexOfFirstBracket, indexOfLastBracket
}