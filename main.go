package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Output)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func Output(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Random number is: %v", randomNumbers())
}

func randomNumbers() int {
	return rand.Intn(1000)
}
