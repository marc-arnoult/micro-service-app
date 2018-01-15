package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayhelloName)     // set router
	err := http.ListenAndServe(":80", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	todo := map[string]interface{}{"name": "marc", "age": 26}
	todoJSON, _ := json.Marshal(todo)

	w.Header().Set("Content-Type", "application/json")
	w.Write(todoJSON)
}
