package main

import (
	"excerices/GoPl/seven"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Data base:")

	// seven.Db.PrintDatabase(os.Stdout)
	// log.Fatal(http.ListenAndServe("localhost:8080", seven.Db))
	log.Fatal(http.ListenAndServe("localhost:8080", seven.Mux))
}
