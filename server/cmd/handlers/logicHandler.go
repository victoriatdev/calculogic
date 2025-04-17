package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"fyp-server/cmd/logic"
	"fyp-server/cmd/utils"
	"fyp-server/collections"
	"net/http"
	"regexp"
	"slices"
	"strings"

	"github.com/adamay909/logicTools/gentzen"
	"github.com/google/uuid"
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
	formulaTree := collections.Node{Id: uuid.New()}
	if err := c.Bind(sequentCalculusFormulae); err != nil {
		return c.String(http.StatusBadRequest, "The provided JSON is invalid.")
	}

	if len(sequentCalculusFormulae.Formula) == 0 {
		return c.String(http.StatusBadRequest, "The provided formula is empty.")
	}

	isProvable, jsonTree, e := AttemptSequentCalculusProof(sequentCalculusFormulae.Formula, &formulaTree)

	if e != nil {
		return c.String(http.StatusInternalServerError, "Server error: "+e.Error())
	}

	if !isProvable {
		return c.String(http.StatusUnprocessableEntity, "The formula could not be proven.")
	}

	// fmt.Println(formulaTree)
	return c.JSONBlob(http.StatusOK, jsonTree)
}

func AttemptSequentCalculusProof2(sequent [][]string, tree *collections.Node) (bool, error) {

	// base case: leaf?
	antecedent := sequent[0]
	succedent := sequent[1]
	fmt.Printf("Current sequent: %v\n", sequent)
	// fmt.Printf("Tree Parent: %v", tree.GetParent())
	// fmt.Printf("Current Tree: %p = %v\n", tree, tree)

	if logic.ApplyAssumption(sequent) {
		fmt.Println("Assumption applied")
		tree.SetRule("A")
		tree.AddChildren([]*collections.Node{})
		return true, nil
	}

	if len(succedent) == 0 {
		return false, errors.New("succedent is empty")
	}

	if len(antecedent) > 0 {
		i := 0
		fmt.Println("Scanning A")

		for i < len(antecedent) {
			if antecedent[i] == logic.Implication {
				fmt.Println("Implication Left detected")
				childSequent1, childSequent2, err := logic.ApplyImplicationLeft(sequent)

				if err != nil {
					return false, err
				}

				// create 2 child nodes for tree
				n1 := collections.Node{Id: uuid.New(), Data: childSequent1, Children: nil}
				n2 := collections.Node{Id: uuid.New(), Data: childSequent2, Children: nil}
				childNodes := append([]*collections.Node{}, &n1, &n2)
				tree.AddChildren(childNodes)
				tree.SetRule("→L")

				p1, e1 := AttemptSequentCalculusProof2(childSequent1, &n1)
				p2, e2 := AttemptSequentCalculusProof2(childSequent2, &n2)
				if e1 != nil {
					return false, e1
				}
				if e2 != nil {
					return false, e2
				}

				// data := tree.GetData().([][]string)
				// formattedAntecedent := utils.ConvertToInfix(data[0])
				// formattedSuccedent := utils.ConvertToInfix(data[1])
				// formattedSequent := append([][]string{}, formattedAntecedent, formattedSuccedent)
				// fmt.Println(formattedSequent)
				// tree.SetData(formattedSequent)

				return p1 && p2, nil
			}

			if antecedent[i] == logic.Conjunction {
				fmt.Println("Conjunction Left detected")
				childSequent, e := logic.ApplyConjunctionLeft(sequent)

				if e != nil {
					return false, e
				}

				node := collections.Node{Id: uuid.New(), Data: childSequent, Children: nil}
				tree.AddChild(&node)
				tree.SetRule("∧L")

				// formattedAntecedent := utils.ConvertToInfix(append([]string{}, sequent[0]...))
				// formattedSuccedent := utils.ConvertToInfix(append([]string{}, sequent[1]...))
				// formattedSequent := append([][]string{}, formattedAntecedent, formattedSuccedent)
				// tree.SetData(formattedSequent)

				p, e := AttemptSequentCalculusProof2(childSequent, &node)
				if e != nil {
					return false, e
				}
				return p, nil
			}

			if antecedent[i] == logic.Disjunction {
				childSequent1, childSequent2 := logic.ApplyDisjunctionLeft(sequent)

				// create 2 child nodes for tree
				n1 := collections.Node{Id: uuid.New(), Data: childSequent1}
				n2 := collections.Node{Id: uuid.New(), Data: childSequent2}
				childNodes := append([]*collections.Node{}, &n1, &n2)
				tree.AddChildren(childNodes)
				tree.SetRule("∨L")

				// formattedAntecedent := utils.ConvertToInfix(append([]string{}, sequent[0]...))
				// formattedSuccedent := utils.ConvertToInfix(append([]string{}, sequent[1]...))
				// formattedSequent := append([][]string{}, formattedAntecedent, formattedSuccedent)
				// tree.SetData(formattedSequent)

				p1, e1 := AttemptSequentCalculusProof2(childSequent1, &n1)
				p2, e2 := AttemptSequentCalculusProof2(childSequent2, &n2)

				if e1 != nil {
					return false, e1
				}
				if e2 != nil {
					return false, e2
				}
				return p1 && p2, nil
			}

			if antecedent[i] == logic.Negation {
				fmt.Println("Negation Left detected")
				childSequent := logic.ApplyNegationLeft(sequent)

				node := collections.Node{Id: uuid.New(), Data: childSequent, Children: nil}
				tree.AddChild(&node)
				tree.SetRule("¬L")

				// formattedAntecedent := utils.ConvertToInfix(append([]string{}, sequent[0]...))
				// formattedSuccedent := utils.ConvertToInfix(append([]string{}, sequent[1]...))
				// formattedSequent := append([][]string{}, formattedAntecedent, formattedSuccedent)
				// tree.SetData(formattedSequent)

				p, e := AttemptSequentCalculusProof2(childSequent, &node)

				if e != nil {
					return false, e
				}

				return p, nil

			}
			i += 1
		}
	}

	// below handles right side cases, need to handle left cases
	// recursive case: there is a still an operator to be attempted
	i := 0
	for i < len(succedent) {
		if succedent[i] == logic.Implication {
			fmt.Println("Implication detected")
			// fmt.Printf("before: %v", sequent)
			childSequent, childSequentError := logic.ApplyImplicationRight(sequent)
			// fmt.Printf("after: %v", sequent)

			if childSequentError != nil {
				return false, childSequentError
			}

			node := collections.Node{Id: uuid.New(), Data: childSequent, Children: nil}
			tree.AddChild(&node)
			tree.SetRule("→R")

			// fmt.Print(sequent)

			// // data := tree.GetData().([][]string)
			// formattedAntecedent := utils.ConvertToInfix(append([]string{}, sequent[0]...))
			// // fmt.Println(formattedAntecedent)
			// formattedSuccedent := utils.ConvertToInfix(append([]string{}, sequent[1]...))
			// formattedSequent := append([][]string{}, formattedAntecedent, formattedSuccedent)
			// // fmt.Println(formattedSequent)
			// tree.SetData(formattedSequent)

			fmt.Printf("before: %v\n", sequent)
			p, e := AttemptSequentCalculusProof2(childSequent, &node)
			if e != nil {
				return false, e
			}
			fmt.Printf("after: %v\n", sequent)

			return p, nil
		}

		if succedent[i] == logic.Conjunction {
			fmt.Println("Conjunction Right detected")
			childSequent1, childSequent2, e := logic.ApplyConjunctionRight(sequent)

			if e != nil {
				return false, e
			}

			n1 := collections.Node{Id: uuid.New(), Data: childSequent1, Children: nil}
			n2 := collections.Node{Id: uuid.New(), Data: childSequent2, Children: nil}
			tree.AddChild(&n1)
			tree.AddChild(&n2)
			tree.SetRule("∧R")

			// formattedAntecedent := utils.ConvertToInfix(append([]string{}, sequent[0]...))
			// formattedSuccedent := utils.ConvertToInfix(append([]string{}, sequent[1]...))
			// formattedSequent := append([][]string{}, formattedAntecedent, formattedSuccedent)
			// tree.SetData(formattedSequent)

			p1, e1 := AttemptSequentCalculusProof2(childSequent1, &n1)
			p2, e2 := AttemptSequentCalculusProof2(childSequent2, &n2)
			if e1 != nil {
				return false, e1
			}
			if e2 != nil {
				return false, e2
			}

			return p1 && p2, nil
		}

		if succedent[i] == logic.Disjunction {
			childSequent := logic.ApplyDisjunctionRight(sequent)

			// create 2 child nodes for tree
			n1 := collections.Node{Id: uuid.New(), Data: childSequent, Children: nil}
			// childNodes := append([]*collections.Node{}, &n1)
			tree.AddChild(&n1)
			tree.SetRule("∨R")

			// formattedAntecedent := utils.ConvertToInfix(append([]string{}, sequent[0]...))
			// formattedSuccedent := utils.ConvertToInfix(append([]string{}, sequent[1]...))
			// formattedSequent := append([][]string{}, formattedAntecedent, formattedSuccedent)
			// tree.SetData(formattedSequent)

			p, e := AttemptSequentCalculusProof2(childSequent, &n1)
			if e != nil {
				return false, e
			}
			return p, nil
		}

		if succedent[i] == logic.Negation {
			fmt.Println("Negation detected")
			childSequent := logic.ApplyNegationRight(sequent)

			node := collections.Node{Id: uuid.New(), Data: childSequent, Children: nil}
			tree.AddChild(&node)
			tree.SetRule("¬R")

			// formattedAntecedent := utils.ConvertToInfix(append([]string{}, sequent[0]...))
			// formattedSuccedent := utils.ConvertToInfix(append([]string{}, sequent[1]...))
			// formattedSequent := append([][]string{}, formattedAntecedent, formattedSuccedent)
			// tree.SetData(formattedSequent)

			p, e := AttemptSequentCalculusProof2(childSequent, &node)

			if e != nil {
				return false, e
			}

			return p, nil

		}

		// fmt.Println("No operator found at: %v", i)
		i += 1
	}

	return false, errors.New("unprovable formula")
}

func AttemptSequentCalculusProof(formula string, tree *collections.Node) (isProvable bool, jsonTree []byte, e error) {

	// fmt.Println(formula)

	formulaList := [][]string{}

	if slices.Contains(strings.Split(formula, ""), logic.Turnstile) {
		splitFormula := strings.Split(formula, "")
		turnstileIndex := slices.Index(splitFormula, logic.Turnstile)
		antecedentFormula, antecedentFormulaError := utils.BuildFormulaTokenList(cleanString(strings.Join(splitFormula[:turnstileIndex], "")))

		if antecedentFormulaError != nil {
			return false, nil, antecedentFormulaError
		}

		succedentFormula, succedentFormulaError := utils.BuildFormulaTokenList(cleanString(strings.Join(splitFormula[turnstileIndex+1:], "")))

		if succedentFormulaError != nil {
			return false, nil, succedentFormulaError
		}

		formulaList = append(formulaList, antecedentFormula, succedentFormula)
	} else {
		formulaListv, e := utils.BuildFormulaTokenList(cleanString(formula))
		// fmt.Printf("%v\n", formulaListv)

		if e != nil {
			return false, nil, e
		}

		formulaList = append([][]string{}, []string{}, formulaListv)
	}

	fmt.Printf("%v\n", formulaList)
	tree.SetData(formulaList)

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

	isProvable, e = AttemptSequentCalculusProof2(formulaList, tree)
	fmt.Printf("isProvable: %v\n", isProvable)
	// fmt.Printf("Final Tree: %p = %v\n", tree, tree)
	// fmt.Printf("Final Tree: %p = %v\n", tree.GetChildren()[0], tree.GetChildren()[0])
	// fmt.Printf("Final Tree: %p = %v\n", tree.GetChildren()[0].GetChildren()[0], tree.GetChildren()[0].GetChildren()[0])
	// fmt.Printf("Final Tree: %p = %v\n", tree.GetChildren()[0].GetChildren()[0].GetChildren()[0], tree.GetChildren()[0].GetChildren()[0].GetChildren()[0])
	// fmt.Printf("Final Tree: %p = %v\n", tree.GetChildren()[0].GetChildren()[0].GetChildren()[1], tree.GetChildren()[0].GetChildren()[0].GetChildren()[1])

	jsonTree, err := json.Marshal(tree)

	if err != nil {
		fmt.Println(err)
		return isProvable, nil, err
	}

	// fmt.Printf("JSON Tree = %v", string(jsonTree))

	return isProvable, jsonTree, e
}

// add support to flip brackets [ ( -> X -> )  ]
func buildFormulaTokenList(s string) (tokenList []string, e error) {
	tokenList = strings.Split(s, "")

	// prefix
	slices.Reverse(tokenList)
	fmt.Println(tokenList)

	reversed := strings.Join(tokenList, "")
	reversed = strings.ReplaceAll(reversed, "(", "%")
	reversed = strings.ReplaceAll(reversed, ")", "(")
	reversed = strings.ReplaceAll(reversed, "%", ")")
	tokenList = strings.Split(reversed, "")

	fmt.Println(tokenList)

	tokenList = convertToPostfix(tokenList)

	fmt.Println(tokenList)

	postfix := strings.Join(tokenList, "")
	postfix = strings.ReplaceAll(postfix, "(", "%")
	postfix = strings.ReplaceAll(postfix, ")", "(")
	postfix = strings.ReplaceAll(postfix, "%", ")")
	tokenList = strings.Split(postfix, "")

	slices.Reverse(tokenList)

	fmt.Println(tokenList)

	return tokenList, e
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
			// fmt.Printf("variable:%s\n", sList[i])
			postfix = append(postfix, sList[i])
			// fmt.Println(postfix)
		} else if sList[i] == "(" {
			stack.Push(sList[i])
			// fmt.Printf("top of stack:%s\n", stack.Top())
			postfix = append(postfix, sList[i])
			// fmt.Printf("stack: %v\n", stack.Inspect())
		} else if sList[i] == ")" {
			for !stack.IsEmpty() && stack.Top() != "(" {
				postfix = append(postfix, stack.Pop().(string))
				postfix = append(postfix, sList[i])
				// fmt.Printf("stack: %v\n", stack.Inspect())
			}
			// fmt.Printf("top of stack:%s\n", stack.Top())
			// postfix = append(postfix, stack.Pop().(string))
			// postfix = append(postfix, stack.Pop().(string))
			stack.Pop()
		} else {
			// fmt.Printf("operator:%s\n", sList[i])
			for !stack.IsEmpty() && (utils.DetermineOperatorPrecedence(sList[i]) <= utils.DetermineOperatorPrecedence(stack.Top().(string))) {
				postfix = append(postfix, stack.Pop().(string))
				// fmt.Println(postfix)
			}
			stack.Push(sList[i])
			// fmt.Printf("top of stack:%s\n", stack.Top())
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
