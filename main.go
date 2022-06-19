package main

import (
	"net/http"

	"github.com/atulanand206/snakes-n-ladders/snip"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// snip.SnakesAndLadders{}.Start()

	routes := snip.Routes()
	handler := http.HandlerFunc(routes.ServeHTTP)
	http.ListenAndServe(":5000", handler)
}
