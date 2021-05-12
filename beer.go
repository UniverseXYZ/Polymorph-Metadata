package functions

import (
	"encoding/json"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type randomBeer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	BrewersTips string `json:"brewers_tips"`
}

type beerFunctionResponse struct {
	APIResponse
	Data randomBeer `json:"data,omitempty"`
}

func RandBeer(w http.ResponseWriter, r *http.Request) {
	log.Printf("Params: %v", r.RequestURI)

	resp, err := http.Get("https://api.punkapi.com/v2/beers/random")
	if err != nil {
		log.Fatalf("response %v\n", err)
		render.JSON(w, r, beerFunctionResponse{APIResponse{false, err.Error()}, randomBeer{}})
		return
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var beers []randomBeer
	err = decoder.Decode(&beers)
	if err != nil {
		log.Fatalf("response %v\n", err)
		render.JSON(w, r, beerFunctionResponse{APIResponse{false, err.Error()}, randomBeer{}})
		return
	}
	log.Printf("%v\n", beers)

	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, beerFunctionResponse{APIResponse{true, ""}, beers[0]})
}
