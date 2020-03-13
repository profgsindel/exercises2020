package main

import (
	"log"
	"net/http"
)

func main() {
	var c http.Client
	resp, err := c.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v: %v", resp.Status, resp.StatusCode)
}
