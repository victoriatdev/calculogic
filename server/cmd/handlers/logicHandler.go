package handlers

import (
	"errors"
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

func ProveSequentCalculus(c echo.Context) error {

	sequentCalculusFormulae := new(SCFormula)
	if err := c.Bind(sequentCalculusFormulae); err != nil {
		return c.String(http.StatusBadRequest, "The provided JSON is invalid.")
	}

	if len(sequentCalculusFormulae.Formula) == 0 {
		return c.String(http.StatusBadRequest, "The provided formula is empty.")
	}

	fmt.Println(sequentCalculusFormulae.Formula)

	isProvable, e := AttemptSequentCalculusProof(sequentCalculusFormulae.Formula)

	if e != nil {
		return c.String(http.StatusInternalServerError, "Server error: "+e.Error())
	}

	if !isProvable {
		return c.String(http.StatusUnprocessableEntity, "The formula could not be proven.")
	}

	return c.String(http.StatusOK, "Formula proven.")
}

func AttemptSequentCalculusProof2(sequent [][]string) (bool, error) {

	// base case: leaf?
	antecedent := sequent[0]
	succedent := sequent[1]
	fmt.Printf("Current sequent: %v\n", sequent)

	if logic.ApplyAssumption(sequent) {
		fmt.Println("Assumption applied")
		// return something???
		return true, nil
	}

	if len(succedent) == 0 {
		return false, errors.New("succedent is empty")
	}

	if len(antecedent) > 0 {
		if antecedent[0] == logic.Implication {
			fmt.Println("Implication Left detected")
			childSequent1, childSequent2, err := logic.ApplyImplicationLeft(sequent)

			if err != nil {
				return false, err
			}

			p1, e1 := AttemptSequentCalculusProof2(childSequent1)
			p2, e2 := AttemptSequentCalculusProof2(childSequent2)
			if e1 != nil {
				return false, e1
			}
			if e2 != nil {
				return false, e2
			}

			return p1 && p2, nil
		}

		if antecedent[0] == logic.Conjunction {
			fmt.Println("Conjunction Left detected")
			childSequent, e := logic.ApplyConjunctionLeft(sequent)

			if e != nil {
				return false, e
			}

			p, e := AttemptSequentCalculusProof2(childSequent)
			if e != nil {
				return false, e
			}
			return p, nil
		}

		if antecedent[0] == logic.Disjunction {
			childSequent1, childSequent2 := logic.ApplyDisjunctionLeft(sequent)
			p1, e1 := AttemptSequentCalculusProof2(childSequent1)
			p2, e2 := AttemptSequentCalculusProof2(childSequent2)

			if e1 != nil {
				return false, e1
			}
			if e2 != nil {
				return false, e2
			}
			return p1 && p2, nil
		}

		if slices.Contains(antecedent, logic.Negation) {
			fmt.Println("Negation Left detected")
			childSequent := logic.ApplyNegationLeft(sequent)
			p, e := AttemptSequentCalculusProof2(childSequent)

			if e != nil {
				return false, e
			}

			return p, nil

		}
	}

	// below handles right side cases, need to handle left cases
	// recursive case: there is a still an operator to be attempted
	if succedent[0] == logic.Implication {
		fmt.Println("Implication detected")
		childSequent, childSequentError := logic.ApplyImplicationRight(sequent)

		if childSequentError != nil {
			return false, childSequentError
		}

		p, e := AttemptSequentCalculusProof2(childSequent)
		if e != nil {
			return false, e
		}
		return p, nil
	}

	if succedent[0] == logic.Conjunction {
		fmt.Println("Conjunction detected")
		childSequent1, childSequent2, e := logic.ApplyConjunctionRight(sequent)

		if e != nil {
			return false, e
		}

		p1, e1 := AttemptSequentCalculusProof2(childSequent1)
		p2, e2 := AttemptSequentCalculusProof2(childSequent2)
		if e1 != nil {
			return false, e1
		}
		if e2 != nil {
			return false, e2
		}
		return p1 && p2, nil
	}

	if succedent[0] == logic.Disjunction {
		childSequent := logic.ApplyDisjunctionRight(sequent)
		p, e := AttemptSequentCalculusProof2(childSequent)
		if e != nil {
			return false, e
		}
		return p, nil
	}

	if slices.Contains(succedent, logic.Negation) {
		fmt.Println("Negation detected")
		childSequent := logic.ApplyNegationRight(sequent)
		p, e := AttemptSequentCalculusProof2(childSequent)

		if e != nil {
			return false, e
		}

		return p, nil

	}

	return false, errors.New("unprovable formula")
}

func AttemptSequentCalculusProof(formula string) (isProvable bool, e error) {

	fmt.Println(formula)

	formulaList := [][]string{}

	if slices.Contains(strings.Split(formula, ""), logic.Turnstile) {
		splitFormula := strings.Split(formula, "")
		turnstileIndex := slices.Index(splitFormula, logic.Turnstile)
		antecedentFormula, antecedentFormulaError := buildFormulaTokenList(cleanString(strings.Join(splitFormula[:turnstileIndex], "")))

		if antecedentFormulaError != nil {
			return false, antecedentFormulaError
		}

		succedentFormula, succedentFormulaError := buildFormulaTokenList(cleanString(strings.Join(splitFormula[turnstileIndex+1:], "")))

		if succedentFormulaError != nil {
			return false, succedentFormulaError
		}

		formulaList = append(formulaList, antecedentFormula, succedentFormula)
	} else {
		formulaListv, e := buildFormulaTokenList(cleanString(formula))
		// fmt.Printf("%v\n", formulaListv)

		if e != nil {
			return false, e
		}

		formulaList = append([][]string{}, []string{}, formulaListv)
	}

	fmt.Printf("%v\n", formulaList)

	// fmt.Printf("%v\n", formulaListv)
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

		now that we have a structure for the sequent (i.e [[A],[S]]) we can abstract the type out to have:
		{
			sequent: [[A],[S]]
			inferenceRule: enum? f
		}

		sequent : Sequent
		sequent.data = [[],[]]

		might help when serialising the data further on for displaying the proof?



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
	// sequent := [][]string{{"A", "∨", "P", "Q"}, {"B"}}
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

	// P -> Q -> P^Q

	// test1, test2 := logic.ApplyConjunctionRight(logic.ApplyImplicationRight(logic.ApplyImplicationRight(sequent)))

	fmt.Printf("Testing on: %v\n", formulaList)

	// fmt.Printf("Left: %v\n", test1)
	// fmt.Printf("Right: %v\n", test2)
	// fmt.Printf("Test on %v: %v and %v\n", sequent, test1, test2)

	// verify that the tree has no false assumptions
	// if not, then isProvable = true
	// else isProvable = false

	isProvable, e = AttemptSequentCalculusProof2(formulaList)
	fmt.Printf("isProvable: %v\n", isProvable)

	return isProvable, e
}

// add support to flip brackets [ ( -> X -> )  ]
func buildFormulaTokenList(s string) (tokenList []string, e error) {
	tokenList = strings.Split(s, "")

	// prefix
	slices.Reverse(tokenList)

	reversed := strings.Join(tokenList, "")
	reversed = strings.ReplaceAll(reversed, "(", "%")
	reversed = strings.ReplaceAll(reversed, ")", "(")
	reversed = strings.ReplaceAll(reversed, "%", ")")
	tokenList = strings.Split(reversed, "")

	fmt.Println(tokenList)

	tokenList = convertToPostfix(tokenList)

	postfix := strings.Join(tokenList, "")
	postfix = strings.ReplaceAll(postfix, "(", "%")
	postfix = strings.ReplaceAll(postfix, ")", "(")
	postfix = strings.ReplaceAll(postfix, "%", ")")
	tokenList = strings.Split(postfix, "")

	slices.Reverse(tokenList)

	fmt.Println(tokenList)

	return tokenList, e
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

// TODO: Add bracket support
// Use (A->B)->A for this
// -(-AB)A
func convertToPostfix(sList []string) []string {
	stack := new(collections.Stack)
	postfix := []string{}
	fmt.Println(sList)

	for i := range len(sList) {
		re := regexp.MustCompile("[A-Z]")

		if re.MatchString(sList[i]) {
			// fmt.Println("variable")
			postfix = append(postfix, sList[i])
			// fmt.Println(postfix)
		} else if sList[i] == "(" {
			stack.Push(sList[i])
			fmt.Println(stack.Top())
		} else if sList[i] == ")" {
			for !stack.IsEmpty() && stack.Top() != "(" {
				postfix = append(postfix, stack.Pop().(string))
			}
			fmt.Println(stack.Top())
			// postfix = append(postfix, stack.Pop().(string))
			stack.Pop()
		} else {
			// fmt.Println("operator")
			for !stack.IsEmpty() && (determineOperatorPrecedence(sList[i]) <= determineOperatorPrecedence(stack.Top().(string))) {
				postfix = append(postfix, stack.Pop().(string))
				// fmt.Println(postfix)
			}
			stack.Push(sList[i])
			fmt.Println(stack.Top())
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
// brackets cant be taken away because we cannot condense expressions inside brackets
// i.e: for |- (p^q) -> p; we cannot condense p^q first without doing the ->R to split it into p^q |- p
func cleanString(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	// s = strings.ReplaceAll(s, "(", ")")
	// s = strings.ReplaceAll(s, ")", "(")

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
