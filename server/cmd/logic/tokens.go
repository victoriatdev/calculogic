package logic

import "strings"

type token struct {
	TokenT tokenId
	String string
}

type tokenList []token

const (
	tConjunction tokenId = 1 << iota
	tDisjunction
	tNegation
	tImplication
	tTerm
)

const (
	tConnective tokenId = tConjunction | tDisjunction | tImplication | tNegation
)

// arity
const (
	tUnary tokenId = tNegation
)

const (
	tBinary tokenId = tConjunction | tDisjunction | tImplication
)

type tokenId int

// clean string by removing any quotes, spaces or tabs
// can also remove brackets because they dont actually affect "precedence"
func cleanString(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")

	return s
}

func toString(t token) string {
	return t.String
}

func createToken(s string) (t token, e error) {

	return t, e
}

func tokenise(s string) (t tokenList, e error) {
	// clean the string
	s = cleanString(s)

	// iterate over the string, creating tokens as we go
	var token token
	for len(s) > 0 {
		// create the token
		token, e = createToken(s)
		t = append(t, token)
	}

	return t, e
}
