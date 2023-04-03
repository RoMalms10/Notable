package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// if I end up using the other folders: import "../handlers"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people []Person

func main() {
	// if I end up using the other folders: home would turn into handlers.HomeHandler (or w/e the function is called)
	http.HandleFunc("/", home)
	http.HandleFunc("/people", getPeople)
	http.HandleFunc("/person", addPerson)
	http.HandleFunc("/delete", deletePerson)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// To test this use: curl -X POST -H "Content-Type: application/json" http://localhost:8080/
// would be overwritten by the function in handlers/home.go
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// To test this use: curl -X POST -H "Content-Type: application/json" -d '{"name": "Alice", "age": 30}' http://localhost:8080/person
func addPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// To test this use: curl -X DELETE -H "Content-Type: application/json" -d '{"name": "Alice", "age": 30}' http://localhost:8080/delete
func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, p := range people {
		if p.Name == person.Name && p.Age == person.Age {
			people = append(people[:i], people[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}
