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
	"time"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var port string

func init() {
	gotenv.Load()
	port = getPort()
	file, _ := ioutil.ReadFile("logo.txt")
	fmt.Println(string(file))
	services.GetClient()
	fmt.Printf("Running on port %v ...\n", port)
}
func main() {

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
	defer services.TimeTrack(time.Now(), "Fibonacci")
	json.NewEncoder(w).Encode(services.Fibonacci(num))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}
	return "8080"
}
