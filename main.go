package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/arminazimi/restapi/handlers"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
