package utils

import (
	"fmt"
	"fyp-server/collections"
	"regexp"
	"slices"
	"strings"
)

func HandleBrackets(inputList []string) (firstIndex int, lastIndex int) {
	indexOfFirstBracket := slices.Index(inputList, "(") + 1
	slices.Reverse(inputList)
	indexOfLastBracket := len(inputList) - slices.Index(inputList, ")") - 1
	slices.Reverse(inputList)
	// bracketedContent := inputList[indexOfFirstBracket+1 : indexOfLastBracket]

	return indexOfFirstBracket, indexOfLastBracket
}

func IsLogicalOperator(s string) bool {
	logicalOperatorList := []string{
		"→", "¬", "∧", "∨", "⊢", /*"∃", "∀",*/
	}
	return slices.Contains(logicalOperatorList, s)
}

func DetermineOperatorPrecedence(operator string) int {
	if operator == "⊢" {
		return 5
	}
	if operator == "¬" {
		return 4
	}
	if operator == "∧" {
		return 3
	}
	if operator == "∨" {
		return 2
	}
	if operator == "→" {
		return 1
	}
	return -1
}

func convertToPostfix(sList []string) []string {
	stack := new(collections.Stack)
	postfix := []string{}
	fmt.Println(sList)

	for i := range len(sList) {
		re := regexp.MustCompile("[A-Z]")

		if re.MatchString(sList[i]) {
			fmt.Printf("variable:%s\n", sList[i])
			postfix = append(postfix, sList[i])
			fmt.Println(postfix)
		} else if sList[i] == "(" {
			stack.Push(sList[i])
			// fmt.Printf("top of stack:%s\n", stack.Top())
			postfix = append(postfix, sList[i])
			// fmt.Printf("stack: %v\n", stack.Inspect())
		} else if sList[i] == ")" {
			// tmp := false
			fmt.Println(stack.IsEmpty())
			for !stack.IsEmpty() && stack.Top() != "(" {
				postfix = append(postfix, stack.Pop().(string))
				postfix = append(postfix, sList[i])
				fmt.Printf("stack: %v\n", stack.Inspect())
			}

			fmt.Println(postfix)
			fmt.Println(stack.Inspect()...)

			if !stack.IsEmpty() && stack.Top() == "(" {
				// postfix = append(postfix, sList[i])
				if postfix[len(postfix)-1] != ")" {
					postfix = append(postfix, sList[i])
				}
			}
			fmt.Printf("top of stack:%s\n", stack.Top())
			// postfix = append(postfix, stack.Pop().(string))
			// postfix = append(postfix, stack.Pop().(string))
			stack.Pop()
		} else {
			fmt.Printf("operator:%s\n", sList[i])
			for !stack.IsEmpty() && (DetermineOperatorPrecedence(sList[i]) <= DetermineOperatorPrecedence(stack.Top().(string))) {
				postfix = append(postfix, stack.Pop().(string))
				fmt.Println(postfix)
			}
			stack.Push(sList[i])
			fmt.Printf("top of stack:%s\n", stack.Top())
		}
	}

	for !stack.IsEmpty() {
		postfix = append(postfix, stack.Pop().(string))
	}

	// fmt.Println(postfix)

	return postfix
}

// add support to flip brackets [ ( -> X -> )  ]
func BuildFormulaTokenList(s string) (tokenList []string, e error) {
	tokenList = strings.Split(s, "")

	// prefix
	slices.Reverse(tokenList)
	// fmt.Println(tokenList)

	reversed := strings.Join(tokenList, "")
	reversed = strings.ReplaceAll(reversed, "(", "%")
	reversed = strings.ReplaceAll(reversed, ")", "(")
	reversed = strings.ReplaceAll(reversed, "%", ")")
	tokenList = strings.Split(reversed, "")

	// fmt.Println(tokenList)

	tokenList = convertToPostfix(tokenList)

	fmt.Println(tokenList)

	postfix := strings.Join(tokenList, "")
	postfix = strings.ReplaceAll(postfix, "(", "%")
	postfix = strings.ReplaceAll(postfix, ")", "(")
	postfix = strings.ReplaceAll(postfix, "%", ")")
	tokenList = strings.Split(postfix, "")

	slices.Reverse(tokenList)

	// fmt.Println(tokenList)

	return tokenList, e
}

func ConvertToInfix(sList []string) []string {
	// fmt.Println(sList)
	slices.Reverse(sList)
	stack := new(collections.Stack)
	infix := ""
	// fmt.Println(sList)

	for i := range len(sList) {
		re := regexp.MustCompile("[A-Z]")

		if re.MatchString(sList[i]) {
			// fmt.Printf("variable:%s\n", sList[i])
			stack.Push(sList[i])
		} else if sList[i] == "(" {
			stack.Push(sList[i])
			// fmt.Printf("top of stack:%s\n", stack.Top())
			// postfix = append(postfix, sList[i])
			// fmt.Printf("stack: %v\n", stack.Inspect())
		} else if sList[i] == ")" {
			// tmp := false
			fmt.Println(stack.IsEmpty())
			for !stack.IsEmpty() && stack.Top() != "(" {
				// postfix = append(postfix, stack.Pop().(string))
				// postfix = append(postfix, sList[i])
				fmt.Printf("stack: %v\n", stack.Inspect())
			}

			// if !stack.IsEmpty() && stack.Top() == "(" {
			// 	// postfix = append(postfix, sList[i])
			// 	if postfix[len(postfix)-1] != ")" {
			// 		postfix = append(postfix, sList[i])
			// 	}
			// }
			// fmt.Printf("top of stack:%s\n", stack.Top())
			// // postfix = append(postfix, stack.Pop().(string))
			// // postfix = append(postfix, stack.Pop().(string))
			// stack.Pop()
		} else {
			// fmt.Printf("operator:%s\n", sList[i])
			// fmt.Printf("stack: %v\n", stack.Inspect()...)
			if len(stack.Inspect()) > 1 && sList[i] != "¬" {
				o2 := stack.Pop().(string)
				o1 := stack.Pop().(string)

				infix = o1 + sList[i] + o2

				stack.Push(infix)
			} else if len(stack.Inspect()) >= 1 && sList[i] == "¬" {
				o1 := stack.Pop().(string)

				infix += o1 + sList[i]

				stack.Push(infix)
			} else {
				infix += sList[i]
			}
			// for !stack.IsEmpty() && (DetermineOperatorPrecedence(sList[i]) <= DetermineOperatorPrecedence(stack.Top().(string))) {
			// 	infix = append(postfix, stack.Pop().(string))
			// 	fmt.Println(postfix)
			// }
			// stack.Push(sList[i])
			// fmt.Printf("top of stack:%s\n", stack.Top())

		}
	}

	infix = Reverse(infix)

	// fmt.Println(infix)

	arr := append([]string{}, infix)

	return arr
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
