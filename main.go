package main

import (
	"fmt"
	"my-bets/bets/infrastructure"
	"net/http"
)

func main() {
	fmt.Println("starting application")
	http.ListenAndServe(":8080", infrastructure.HandlersFactory())
}
