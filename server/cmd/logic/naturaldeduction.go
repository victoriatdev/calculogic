package logic

import (
	"encoding/json"
	"fmt"
	"fyp-server/cmd/utils"
	"fyp-server/collections"
	"net/http"
	"regexp"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type NDFormula struct {
	Formula string `json:"formula"`
}

func ProveNaturalDeduction(c echo.Context) error {

	NDFormula := new(NDFormula)
	formulaTree := collections.Node{Id: uuid.New()}
	if err := c.Bind(NDFormula); err != nil {
		return c.String(http.StatusBadRequest, "The provided JSON is invalid.")
	}

	if len(NDFormula.Formula) == 0 {
		return c.String(http.StatusBadRequest, "The provided formula is empty.")
	}

	// fmt.Println(sequentCalculusFormulae.Formula)

	isProvable, jsonTree, e := AttemptNaturalDeductionProof(NDFormula.Formula, &formulaTree)

	if e != nil {
		return c.String(http.StatusInternalServerError, "Server error: "+e.Error())
	}

	if !isProvable {
		return c.String(http.StatusUnprocessableEntity, "The formula could not be proven.")
	}

	return c.JSONBlob(http.StatusOK, jsonTree)
}

func AttemptNaturalDeductionProof(formula string, tree *collections.Node) (isProvable bool, jsonTree []byte, e error) {

	formulaList := [][]string{}

	formulaListv, e := buildFormulaTokenList(cleanString(formula))

	if e != nil {
		return false, nil, e
	}

	formulaList = append([][]string{}, []string{}, formulaListv)

	fmt.Printf("%v\n", formulaList)
	tree.SetData(formulaList)

	fmt.Printf("Testing on: %v\n", formulaList)

	isProvable, e = AttemptNaturalDeductionProof2(formulaList, tree)

	fmt.Printf("isProvable: %v\n", isProvable)

	jsonTree, err := json.Marshal(tree)

	if err != nil {
		fmt.Println(err)
		return isProvable, nil, err
	}

	fmt.Printf("JSON Tree = %v", string(jsonTree))

	return isProvable, jsonTree, e
}

func AttemptNaturalDeductionProof2(formula [][]string, tree *collections.Node) (isProvable bool, error error) {
	return true, nil
}

func buildFormulaTokenList(s string) (tokenList []string, e error) {
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

	// fmt.Println(tokenList)

	postfix := strings.Join(tokenList, "")
	postfix = strings.ReplaceAll(postfix, "(", "%")
	postfix = strings.ReplaceAll(postfix, ")", "(")
	postfix = strings.ReplaceAll(postfix, "%", ")")
	tokenList = strings.Split(postfix, "")

	slices.Reverse(tokenList)

	// fmt.Println(tokenList)

	return tokenList, e
}

func NaturalDeductionRoot(c echo.Context) error {

	return c.String(http.StatusOK, "Root natural deduction route")
}

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
