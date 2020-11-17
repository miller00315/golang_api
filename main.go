package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Api iniciada")

	router := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", router))
}
