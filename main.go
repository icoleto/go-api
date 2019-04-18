package main

import (
	"encoding/json"
	"fmt"
	"go-api/services"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	port := getPort()
	file, _ := ioutil.ReadFile("logo.txt")
	fmt.Println(string(file))
	fmt.Printf("Running on port %v ...", port)

	router := mux.NewRouter()
	router.HandleFunc("/", hello).Methods("GET")
	router.HandleFunc("/fibonacci/{num}", fibonacciHandler).Methods("GET")
	// router.HandleFunc("/users/{id}").Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go World!")
}
func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num, _ := strconv.Atoi(vars["num"])
	json.NewEncoder(w).Encode(services.Fibonacci(num))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}
	return "8080"
}
