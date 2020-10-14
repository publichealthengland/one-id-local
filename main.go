package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		panic(errors.New("package names are required as args"))
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		js, err := json.Marshal(args[1:])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
