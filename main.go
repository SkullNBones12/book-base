package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

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

	data := string(responseData)

	makeNewFile()
	writeFile(data)
	parseAuthor()
	cleanDoc()

}

func parseAuthor() {
	cmd := exec.Command("./run_script")
	if err := cmd.Run(); err != nil {
		log.Println("Couldn't execute command: ", err)
	}
}

func cleanDoc() {
	cmd := exec.Command("./clean_whitespace")
	if err := cmd.Run(); err != nil {
		log.Println("Couldn't clean whitespace: ", err)
	}
}

func makeNewFile() {
	file, err := os.Create("json")
	if err != nil {
		log.Println("Couldn't create file: ", err)
	}
	defer file.Close()

}

func writeFile(x string) {
	file, err := os.OpenFile("json", os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Couldn't open file: ", err)
	}
	defer file.Close()

	if _, err = file.WriteString(x); err != nil {
		log.Println("Couldn't write to file: ", err)
	}
}
