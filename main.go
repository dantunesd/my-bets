package main

import (
	"my-bets/bets/infrastructure"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", infrastructure.ServerFactory())
}
