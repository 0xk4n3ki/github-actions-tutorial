package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Add(a, b int) int {
	return a+b
}

func additionHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	aStr := values.Get("a")
	bStr := values.Get("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		http.Error(w, "parameter 'a' must be an integer", http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		http.Error(w, "parameter 'a' must be an integer", http.StatusBadRequest)
		return
	}

	result := Add(a, b)
	fmt.Fprintf(w, "%d", result)
}


func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/addition", additionHandler)
	fmt.Println("server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}


func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "visit endpoint like /addition?a=4&b=6")
}