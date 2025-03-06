package handlers

import (
	"fmt"
	"net/http"

	"github.com/adamay909/logicTools/gentzen"
	"github.com/labstack/echo/v4"
)

type Proof struct {
	Proof []string `json:"proof"`
	//Offset int `json:"offset"`
}

func LogicRoot(c echo.Context) error {

	b := []string { "p;p;a", "q;q;a", "p,q;Kpq;1,2,ki", "p,q;K(Kpq)p;1,3,ki" }

				//   [p;p;A q;q;A p,q;Kpq;1,2,ki]
	gentzen.SetDebug(true)
	//gentzen.SetStandardPolish(false)
	g := gentzen.CheckDeriv(b, 1)
	
	// fmt.Print(gentzen.ShowLog())
	fmt.Print(gentzen.PrintDerivation(b,1))
	// fmt.Print(gentzen.ShowDebugLog())
	fmt.Print(g)

	return c.String(http.StatusOK, "Root logic handler route")
}

func HandleLogicProof(c echo.Context) error {

	p := new(Proof)
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	g := gentzen.CheckDeriv(p.Proof, 1);
	if !g {
		fmt.Print(g)
		return c.String(http.StatusOK, "proof invalid")
	}

	return c.String(http.StatusOK, "Proof valid")
}