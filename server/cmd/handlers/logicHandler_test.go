package handlers

import (
	"fyp-server/collections"
	"testing"

	"github.com/google/uuid"
)

// Tests P->Q->P^Q
func TestPPQTestCase(t *testing.T) {
	initialSequent := "P→Q→P∧Q"
	initialTree := collections.Node{Id: uuid.New()}
	isProvable, jsonTree, err := AttemptSequentCalculusProof(initialSequent, &initialTree)

	want := true

	if isProvable != want {
		t.Errorf("PPQ Test case failed, got isProvable = %v, jsonTree = %v, err = %v", isProvable, jsonTree, err)
	}
}


