package handlers

import (
	"github.com/go-chi/render"
	// log "github.com/sirupsen/logrus"
	"math/big"
	"net/http"
)

type APIResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error,omitempty"`
}

type ExampleDoubleRequest struct {
	Number *big.Int
}

type exampleResponse struct {
	APIResponse
	Double string `json:"double"`
}

func (r *ExampleDoubleRequest) double() *big.Int {
	return r.Number.Mul(r.Number, big.NewInt(2))
}

func DoubleNumber(w http.ResponseWriter, r *http.Request) {
	numStr := r.URL.Query().Get("num")
	i := new(big.Int)
	_, ok := i.SetString(numStr, 10)
	if !ok {
		render.JSON(w, r, APIResponse{Status: false, Error: "Could not parse int"})
		return
	}

	exampleRequest := &ExampleDoubleRequest{i}

	double := exampleRequest.double()

	render.JSON(w, r, exampleResponse{APIResponse{Status: true, Error: ""}, double.String()})
}
