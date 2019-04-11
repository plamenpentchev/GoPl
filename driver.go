package main

import (
	"excerises/GoPl/seven"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Data base:")

	// seven.Db.PrintDatabase(os.Stdout)
	// log.Fatal(http.ListenAndServe("localhost:8080", seven.Db))
	go http.ListenAndServe("localhost:8080", seven.Mux)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
