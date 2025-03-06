package logic

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type InferenceRuleEnum int

const (
  ASSUMPTION InferenceRuleEnum = iota
  NEGATION_ELIMINATION
  NEGATION_INTRODUCTION
  DISJUNCTION_ELIMINATION
  DISJUNCTION_INTRODUCTION
  CONJUNCTION_ELIMINATION
  CONJUNCTION_INTRODUCTION
  CONDITIONAL_ELIMINATION
  CONDITIONAL_INTRODUCTION
)

var InferenceRule = map[InferenceRuleEnum]string{
	ASSUMPTION: "a",
	NEGATION_ELIMINATION: "ne",
	NEGATION_INTRODUCTION: "ni",
  	DISJUNCTION_ELIMINATION: "de",
  	DISJUNCTION_INTRODUCTION: "di",
  	CONJUNCTION_ELIMINATION: "ke",
  	CONJUNCTION_INTRODUCTION: "ki",
  	CONDITIONAL_ELIMINATION: "ce",
  	CONDITIONAL_INTRODUCTION: "ci",
}

func (ir InferenceRuleEnum) String() string {
	return InferenceRule[ir]
}

type Sequent struct {
	Antecedent string `json:"Antecedent"`
	Succedent string `json:"Succedent"`
	InferenceRule string `json:"InferenceRule"`
}

type ProofRequest struct {
	Id string `json:"id"`
	Proof []ProofRequest `json:"proof"`
	Sequent Sequent
}

func deconstructProof(str string) (strlist []string) {
	tmp := strings.Split(str, "⊢")
	for _,v := range tmp {
		strlist = append(strlist, strings.TrimSpace(v))
	}
	return strlist
}

func NaturalDeductionRoot(c echo.Context) error {
	
	proof := new(ProofRequest)
	if err := c.Bind(proof); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	fmt.Println(proof.Proof)

	// d := deconstructProof(proof.Proof)
	// fmt.Println(d)

	return c.String(http.StatusOK, "Root natural deduction route")
}