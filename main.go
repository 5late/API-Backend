package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	BirthDate string `json:"Birthday"`
	Age       string `json:"age"`
}

func whereyoulooking(response http.ResponseWriter, request *http.Request) {
	result := `{"status":404,"message":"where are you looking"}`

	var finalResult map[string]interface{}
	json.Unmarshal([]byte(result), &finalResult)

	json.NewEncoder(response).Encode(finalResult)
}

func GetAllPeople(response http.ResponseWriter, request *http.Request) {

}

func main() {
	fmt.Println("Starting the api....")
	route := mux.NewRouter()
	router := cors.Default().Handler(route)

	//route.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	route.HandleFunc("/people/", GetAllPeople).Methods("GET")
	//route.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	//route.HandleFunc("/rmperson/{id}", DeletePersonEndpoint).Methods("DELETE")
	//route.HandleFunc("/changeperson/{id}/{type}/{value}", UpdatePersonEndpoint).Methods("PATCH")
	route.NotFoundHandler = http.HandlerFunc(whereyoulooking)
	http.ListenAndServe(":12345", router)
}
