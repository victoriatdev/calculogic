package logic

func reverseInput(input string) (reversedString string) {
	for _,value := range input {
		reversedString = string(value) + reversedString
	}
	return
}

func prefixTranslator(input string) string {
	reversed := reverseInput(input)
	// postfixEquation, err := postfix.ParseInfix(reversed)
	// eqMap := postfix.Evaluate(postfixEquation)
	// formattedInput := reverseInput(postfix)

	return reversed
}