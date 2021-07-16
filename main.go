package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type People struct {
	People []Person `json:"people"`
}

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	BirthDate string `json:"birthdate"`
	Age       int    `json:"age"`
}

func whereyoulooking(response http.ResponseWriter, request *http.Request) {
	result := `{"status":404,"message":"where are you looking"}`

	var finalResult map[string]interface{}
	json.Unmarshal([]byte(result), &finalResult)

	json.NewEncoder(response).Encode(finalResult)
}

func GetAllPeople(response http.ResponseWriter, request *http.Request) {
	files, _ := ioutil.ReadDir("./people/")
	file_count := len(files)
	for i := 0; i < file_count; i++ {
		json_file, err := os.Open("./people/" + fmt.Sprint(i) + ".json")
		if err != nil {
			log.Fatal(err)
		}
		defer json_file.Close()

		byteValue, _ := ioutil.ReadAll(json_file)

		var people People

		json.Unmarshal(byteValue, &people)
		fmt.Println(people)
		for ii := 0; ii < len(people.People); ii++ {
			fmt.Println("User ID: ", people.People[ii].ID)
			fmt.Println("User First Name: " + people.People[ii].FirstName)
			fmt.Println("User Last Name: " + people.People[ii].LastName)
			fmt.Println("User Birth Date" + people.People[ii].BirthDate)
			fmt.Println("User Age:", people.People[ii].Age)
		}
		json.NewEncoder(response).Encode(people)
	}
}

func main() {
	fmt.Println("Starting the api....")
	route := mux.NewRouter()
	router := cors.Default().Handler(route)

	//route.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	route.HandleFunc("/people", GetAllPeople).Methods("GET")
	//route.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	//route.HandleFunc("/rmperson/{id}", DeletePersonEndpoint).Methods("DELETE")
	//route.HandleFunc("/changeperson/{id}/{type}/{value}", UpdatePersonEndpoint).Methods("PATCH")
	route.NotFoundHandler = http.HandlerFunc(whereyoulooking)
	http.ListenAndServe(":12345", router)
}
