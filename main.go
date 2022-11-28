package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/packages/handlers"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/users/", handlers.UsersRouter)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// sudo lsof -t -i tcp:11111 | sudo xargs kill
