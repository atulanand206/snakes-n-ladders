package snip

import (
	"encoding/json"
	"fmt"
	"net/http"

	net "github.com/atulanand206/go-network"
)

type CreateGameRequest struct {
	Size    int `json:"size"`
	Players int `json:"players"`
}

var (
	game = SnakesAndLadders{}
)

const (
	Err_RequestNotDecoded = "request can't be decoded"
)

func Routes() *http.ServeMux {
	router := http.NewServeMux()
	chain := net.MiddlewareChain{
		net.ApplicationJsonInterceptor(),
		// net.AuthenticationInterceptor(),
	}

	// Interceptor chain with only PUT method.
	// getChain := chain.Add(net.CorsInterceptor(http.MethodGet))
	// putChain := chain.Add(net.CorsInterceptor(http.MethodPut))
	// Interceptor chain with only POST method.
	postChain := chain.Add(net.CorsInterceptor(http.MethodPost))

	router.HandleFunc("/new", postChain.Handler(HandlerNewGame))
	return router
}

func HandlerNewGame(w http.ResponseWriter, r *http.Request) {
	var request CreateGameRequest
	if err := DecodeJsonBody(r, &request); err != nil {
		http.Error(w, Err_RequestNotDecoded, http.StatusInternalServerError)
		return
	}
	board := game.generateGame(request.Size, request.Players)
	fmt.Println(board)
	json.NewEncoder(w).Encode(board)
}

func DecodeJsonBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}
