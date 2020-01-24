package main

import (
	"encoding/json"
	"fmt"
	"go-api/dtos"
	"go-api/services"
	"go-api/use_cases"
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
	fmt.Printf("Running on port %v ...\n", port)
}
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", hello).Methods("GET")
	router.HandleFunc("/favicon.ico", faviconHandler)
	router.HandleFunc("/fibonacci/{num}", fibonacciHandler).Methods("GET")
	router.HandleFunc("/user", getUserByName).Methods("GET").Queries("name", "{name}")

	log.Fatal(http.ListenAndServe(":"+port, router))

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go World!")
}
func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num, _ := strconv.ParseFloat(vars["num"], 64)
	defer services.TimeTrack(time.Now(), "Fibonacci")

	result := &dtos.FibonacciDto{N: num, Value: services.Fibonacci(num)}
	json.NewEncoder(w).Encode(result)
}

func getUserByName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	json.NewEncoder(w).Encode(use_cases.FindByName(name))
}
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}
	return "8080"
}
