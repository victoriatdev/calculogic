package logic

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/adamay909/logicTools/gentzen"
	"github.com/labstack/echo/v4"
)

// Constraints:
// Assume an input of a string in the form <Antecedents> |- <Succedents>
// Succedent cannot be empty (as it implies no goal)
// Antecedent CAN be empty (as it implies full implication)
// Tokenise the string and create an Expression tree
// ex1: P→Q→P∧Q => →P→Q∧PQ
// ex2: P→P∨Q => →P∨PQ
// e

func buildProofSegment(s Sequent) string {
	return s.Antecedent + ";" + s.Succedent + ";" + s.InferenceRule
}

func buildProofSegment2(s Sequent) string {
	return s.Antecedent + ";" + s.Succedent + "," + s.InferenceRule
}

type SCFormula interface {
}

func buildProof(p ProofRequest) []string {

	var proof []string
	for _, e := range p.Proof {
		t := buildProofSegment(e.Sequent)
		proof = append(proof, t)
	}
	proof = append(proof, buildProofSegment2(p.Sequent))

	return proof
}

func SequentCalculusRoot(c echo.Context) error {

	proof := new(ProofRequest)
	// if err := c.Bind(proof); err != nil {
	// 	return c.String(http.StatusBadRequest, "bad request")
	// }
	c.Bind(proof)
	if proof.Sequent.InferenceRule == "∧I" {
		proof.Sequent.InferenceRule = "ki"
	}

	for _, e := range proof.Proof {
		fmt.Println(e.Sequent)
		if e.Sequent.InferenceRule == "A" {
			e.Sequent.InferenceRule = strings.ToLower(e.Sequent.InferenceRule)
		}
	}

	fmt.Println(proof.Sequent)
	// fmt.Println(proof.Id)
	// fmt.Println(proof.Proof)

	d := buildProof(*proof)
	fmt.Println(d)

	// buildProof(*proof)
	// fmt.Println(buildProof(*proof))

	// d := buildProof(*proof)
	// d := strings.ReplaceAll(proof.Proof, "⊢", ";")
	// l := []string {}
	// l = append(l, d)

	// gentzen.SetStandardPolish(false)
	gentzen.SetDebug(true)
	fmt.Println(gentzen.ShowDebugLog())
	fmt.Println(gentzen.ShowLog())
	g := gentzen.CheckDeriv(d, 1)
	// fmt.Println(d)
	fmt.Println(g)

	return c.String(http.StatusOK, "Root natural deduction route")
}


