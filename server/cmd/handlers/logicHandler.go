package handlers

import (
	"fmt"
	"fyp-server/cmd/logic"
	"fyp-server/collections"
	"net/http"
	"regexp"
	"slices"
	"strings"

	"github.com/adamay909/logicTools/gentzen"
	"github.com/labstack/echo/v4"
)

type Proof struct {
	Proof []string `json:"proof"`
	//Offset int `json:"offset"`
}

type SCFormula struct {
	Formula string `json:"formula"`
}

func LogicRoot(c echo.Context) error {

	// b := []string{"p;p;a", "q;q;a", "p,q;Kpq;1,2,ki", "p,q;K(Kpq)p;1,3,ki"}
	// //   [p;p;A q;q;A p,q;Kpq;1,2,ki]
	// gentzen.SetDebug(true)
	// //gentzen.SetStandardPolish(false)
	// g := gentzen.CheckDeriv(b, 1)

	// fmt.Print(gentzen.ShowLog())
	// // fmt.Print(gentzen.PrintDerivation(b,1))
	// fmt.Print(gentzen.ShowDebugLog())
	// fmt.Print(g)

	return c.String(http.StatusOK, "Root logic handler route")
}

// DONE
// func applyAssumption(sequent [][]string) bool {

// 	// check left and right arrays for a common singular element without a logical operator
// 	var antecedent = sequent[0]
// 	var succedent = sequent[1]

// 	if slices.Contains(antecedent, "→") || slices.Contains(antecedent, "¬") ||
// 		slices.Contains(antecedent, "∧") || slices.Contains(antecedent, "∨") {
// 		return false
// 	}

// 	if slices.Contains(succedent, "→") || slices.Contains(succedent, "¬") ||
// 		slices.Contains(succedent, "∧") || slices.Contains(succedent, "∨") {
// 		return false
// 	}

// 	for i := range len(antecedent) { //[Q,P]
// 		if slices.Contains(succedent, antecedent[i]) {
// 			return true
// 		}
// 	}

// 	return false
// }

// IMPLICATION RULES
func applyImplicationRight(sequent [][]string) [][]string {

	var antecedent = sequent[0]
	var succedent = sequent[1]

	// determine succedent values
	// [[], [->, P, P]]

	if !slices.Contains(succedent, "→") {
		return sequent
	}

	implicationOperator := slices.Index(succedent, "→")
	leftMost := succedent[implicationOperator+1]
	rightMost := succedent[implicationOperator+2:]

	// prepend leftmost to antecedent via composite literal

	antecedent = append([]string{leftMost}, antecedent...)

	sequent = append([][]string{}, antecedent, rightMost)

	return sequent
}

func applyImplicationLeft(sequent [][]string) ([][]string, [][]string) {

	antecedent := sequent[0]
	succedent := sequent[1]

	if !slices.Contains(antecedent, "→") {
		return sequent, [][]string{}
	}

	implicationOperator := slices.Index(antecedent, "→")
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

	return leftSequent, rightSequent
}

// NEGATION RULES
func applyNegationRight(sequent [][]string) [][]string {

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

func applyNegationLeft(sequent [][]string) [][]string {

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

// CONJUNCTION RULES
func applyConjunctionRight(sequent [][]string) ([][]string, [][]string) {

	gamma := sequent[0]
	succedent := sequent[1]

	if !slices.Contains(succedent, "∧") {
		return sequent, [][]string{}
	}

	conjunctionOperator := slices.Index(succedent, "∧")
	leftMost := succedent[conjunctionOperator+1]
	rightMost := succedent[conjunctionOperator+2]
	delta := succedent[conjunctionOperator+3:]

	// build leftSequent

	leftSuccedent := append([]string{leftMost}, delta...)
	leftSequent := append([][]string{}, gamma, leftSuccedent)

	// build rightSequent

	rightSuccedent := append([]string{rightMost}, delta...)
	rightSequent := append([][]string{}, gamma, rightSuccedent)

	return leftSequent, rightSequent
}

func applyConjunctionLeft(sequent [][]string) [][]string {

	antecedent := sequent[0]
	delta := sequent[1]

	if !slices.Contains(antecedent, "∧") {
		return sequent
	}

	conjunctionOperator := slices.Index(antecedent, "∧")

	tmpAntecedent := append([]string{}, antecedent[:conjunctionOperator]...)
	fmt.Printf("%v\n", tmpAntecedent)
	tmpAntecedent = append(tmpAntecedent, antecedent[conjunctionOperator+1:]...)

	sequent = append([][]string{}, tmpAntecedent, delta)

	return sequent
}

// DISJUNCTION LEFT
func applyDisjunctionLeft(sequent [][]string) ([][]string, [][]string) {

	antecedent := sequent[0]
	delta := sequent[1]

	if !slices.Contains(antecedent, "∨") {
		return sequent, [][]string{}
	}

	disjunctionOperator := slices.Index(antecedent, "∨")
	leftMost := antecedent[disjunctionOperator+1]
	rightMost := antecedent[disjunctionOperator+2]
	gamma := antecedent[:disjunctionOperator]

	// build leftSequent
	leftAntecedent := append(append([]string{}, gamma...), leftMost)
	leftSequent := append([][]string{}, leftAntecedent, delta)

	// build rightSequent

	rightAntecedent := append(append([]string{}, gamma...), rightMost)
	rightSequent := append([][]string{}, rightAntecedent, delta)

	return leftSequent, rightSequent
}

func applyDisjunctionRight(sequent [][]string) [][]string {
	gamma := sequent[0]
	succedent := sequent[1]

	if !slices.Contains(succedent, "∨") {
		return sequent
	}

	conjunctionOperator := slices.Index(succedent, "∨")

	tmpSuccedent := append([]string{}, succedent[:conjunctionOperator]...)

	tmpSuccedent = append(tmpSuccedent, succedent[conjunctionOperator+1:]...)

	sequent = append([][]string{}, gamma, tmpSuccedent)

	return sequent
}

func ProveSequentCalculus(c echo.Context) error {

	sequentCalculusFormulae := new(SCFormula)
	if err := c.Bind(sequentCalculusFormulae); err != nil {
		return c.String(http.StatusBadRequest, "The provided JSON is invalid.")
	}

	if len(sequentCalculusFormulae.Formula) == 0 {
		return c.String(http.StatusBadRequest, "The provided formula is empty.")
	}

	isProvable, e := AttemptSequentCalculusProof(sequentCalculusFormulae.Formula)

	if !isProvable {
		return c.String(http.StatusUnprocessableEntity, "The formula could not be proven"+e.Error())
	}

	return c.String(http.StatusAccepted, "Formula proven.")
}

func AttemptSequentCalculusProof(formula string) (isProvable bool, e error) {

	formulaListv, e := buildFormulaTokenList(cleanString(formula))

	fmt.Printf("%v\n", formulaListv)
	// formulaList := append([][]string{}, []string{}, formulaListv)

	// now that we have the string list, need to prefix it? DONE
	// identifies "most important operator" for the sequent? (i.e ->EE)
	// problem is how to determine when one sequent ends? in prefix can we "look ahead" in the list
	// 		to find the next operator
	// i.e A -> A ^ B => ->A^AB
	// we match "->" to get "->EE" and assign s1 := [A] and look ahead to find "^" and assign s2 := [^,a,b]
	/*

		important thing to note is what kinds of structure the tokenList will be in. We can have:
		[->,A,^,A,B]
		[A,⊢,^,A,B]
			in this case, we can split the by ⊢ and mandate that ⊢ has the highest priority?
			i dont think that splitting by ⊢ is a good idea - why not just have [[],[]] where the comma acts as the split?
			i.e: [[A],[^,A,B]]?
			then we can pass the whole array in and have access to the antecedent if we need it (i.e ->I)
			[[],[->,A,^,A,B]]
			iterate through left most, then say "nothing is there" hence "right most", then apply rules on right most
			match on "->", pass in the array, and calculate the changes to the arrays:
				i: [[], [->,A,^,A,B]]
				o: [[A], [^,A,B]]
			left mmost has nothing so attempt hypthoetsis but there's no free terms in the right (i.e there are operators)
			hence we go on succ and attempt something
				i: [[A], [^,A,B]]
				o: [[A,]]



		N1 [[],[->,P,->,Q,->,^,P,Q]] prefix conversion only happens once for succedent
		N2 [[P],[->,Q,^,P,Q]]
		N3 [[Q,P],[^,P,Q]]
		N4 [[Q,P],[P]] => no children, hence leaf => for proof to be valid must be A
		N5 [[Q,P],[Q]] => no children, hence leaf => for prove to be valid but be A


	*/
	// seqNode(->, seqNode(a, nil, nil), seqNode(^, seqnode(a, nil, nil), seqNode(b, nil, nil))
	// //tokenTree := logic.Tree()

	// // sequent := [][]string{{"P", "Q"}, {"P"}}
	// // sequent := [][]string{{"A"}, {"B"}}
	// sequent := [][]string{{"P"}, {"→", "Q", "∧", "P", "Q"}}
	// sequent := [][]string{{"A"}, {"∧", "P", "Q", "B"}}
	// sequent := [][]string{{}, {"→", "P", "→", "Q", "∧", "P", "Q"}}
	// sequent := [][]string{{"Q", "P"}, {"∧", "P", "Q", "R"}}
	sequent := [][]string{{"A", "∨", "P", "Q"}, {"B"}}
	// sequent := [][]string{{}, {"¬", "A", "A"}}
	// sequent := [][]string{{"P", "¬", "Q"}, {"P"}}
	// sequent := [][]string{{"A", "∧", "A", "B"}, {"A"}}
	// sequent := [][]string{{"A"}, {"∨", "A", "B", "¬", "A"}}
	// sequent := [][]string{{"P", "R", "→", "A", "B"}, {"Q"}}
	// test := applyAssumption(sequent)
	// test := applyImplicationRight(sequent)
	// test := applyNegationRight(sequent)
	// test := applyNegationLeft(sequent)
	// test := applyDisjunctionRight(sequent)
	// test1, test2 := applyConjunctionRight(sequent)
	// test1v, test2v := applyConjunctionRight(applyImplicationRight(applyImplicationRight(formulaList)))
	// test1 := applyAssumption(test1v)
	// test2 := applyAssumption(test2v)

	fmt.Printf("Testing on: %v\n", sequent)

	logic.ApplyAssumption(sequent)

	test1, test2 := applyDisjunctionLeft(sequent)

	fmt.Printf("Left: %v\n", test1)
	fmt.Printf("Right: %v\n", test2)
	// fmt.Printf("Test on %v: %v and %v\n", sequent, test1, test2)

	// verify that the tree has no false assumptions
	// if not, then isProvable = true
	// else isProvable = false

	isProvable = true

	return isProvable, e
}

func buildFormulaTokenList(s string) (tokenList []string, e error) {
	tokenList = strings.Split(s, "")

	// prefix
	slices.Reverse(tokenList)
	tokenList = convertToPostfix(tokenList)
	slices.Reverse(tokenList)

	for i := range len(tokenList) {
		if !isLogicalOperator(tokenList[i]) {
			// This is saying there is not
		}
	}

	return
}

func isLogicalOperator(s string) bool {
	logicalOperatorList := []string{
		"→", "¬", "∧", "∨", "⊢", /*"∃", "∀",*/
	}
	return slices.Contains(logicalOperatorList, s)
}

func determineOperatorPrecedence(operator string) int {
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

	for i := range len(sList) {
		re := regexp.MustCompile("[A-Z]")

		if re.MatchString(sList[i]) {
			// fmt.Println("variable")
			postfix = append(postfix, sList[i])
			// fmt.Println(postfix)
		} else if sList[i] == "(" {
			stack.Push(sList[i])
		} else if sList[i] == ")" {
			for !stack.IsEmpty() && stack.Top() != "(" {
				postfix = append(postfix, stack.Pop().(string))
			}
			stack.Pop()
		} else {
			// fmt.Println("operator")
			for !stack.IsEmpty() && (determineOperatorPrecedence(sList[i]) <= determineOperatorPrecedence(stack.Top().(string))) {
				postfix = append(postfix, stack.Pop().(string))
				// fmt.Println(postfix)
			}
			stack.Push(sList[i])
		}
	}

	for !stack.IsEmpty() {
		postfix = append(postfix, stack.Pop().(string))
	}

	// fmt.Println(postfix)

	return postfix
}

// clean string by removing any quotes, spaces or tabs
// can also remove brackets because they dont actually affect "precedence"
func cleanString(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")

	return s
}

func HandleLogicProof(c echo.Context) error {

	p := new(Proof)
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	g := gentzen.CheckDeriv(p.Proof, 1)
	if !g {
		fmt.Print(g)
		return c.String(http.StatusOK, "proof invalid")
	}

	return c.String(http.StatusOK, "Proof valid")
}
