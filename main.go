package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Docs struct {
	Arrays []any `json:"docs"`
}

type Name struct {
	Names []any `json:"author_name"`
}

func main() {
	response, err := http.Get("https://openlibrary.org/search.json?title=the+lord+of+the+rings")

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(responseData))
}
