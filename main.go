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
	ID        json.Number `json:"id"`
	FirstName string      `json:"firstname"`
	LastName  string      `json:"lastname"`
	BirthDate string      `json:"birthdate"`
	Age       json.Number `json:"age"`
	DiscordID json.Number `json:"discordid"`
}

type Appointment struct {
	ID        json.Number `json:"id"`
	FirstName string      `json:"firstname"`
	LastName  string      `json:"lastname"`
	Date      string      `json:"date"`
	Time      string      `json:"time"`
	Reason    string      `json:"reason"`
	DiscordID string      `json:"discordid"`
}

func whereyoulooking(response http.ResponseWriter, request *http.Request) {
	result := `{"status":404,"message":"where are you looking"}`

	var finalResult map[string]interface{}
	json.Unmarshal([]byte(result), &finalResult)

	json.NewEncoder(response).Encode(finalResult)
}

func GetAllPeople(response http.ResponseWriter, request *http.Request) {
	json_file, err := os.Open("./people/all.json")
	if err != nil {
		log.Fatal(err)
	}
	defer json_file.Close()

	byteValue, _ := ioutil.ReadAll(json_file)

	var person []Person

	json.Unmarshal(byteValue, &person)
	fmt.Println(person)
	json.NewEncoder(response).Encode(person)
}

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	var person Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person)

	file, _ := json.MarshalIndent(person, "", "    ")

	_ = ioutil.WriteFile("./people/"+fmt.Sprint(person.ID)+".json", file, 0644)

	allfile, err := ioutil.ReadFile("./people/all.json")
	if err != nil {
		log.Println(err)
	}

	datas := []Person{}

	json.Unmarshal(allfile, &datas)

	//Define what we want to add
	newStruct := &Person{
		ID:        person.ID,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		BirthDate: person.BirthDate,
		Age:       person.Age,
		DiscordID: person.DiscordID,
	}

	datas = append(datas, *newStruct)

	//JSON-lize the data defined above
	dataBytes, err := json.MarshalIndent(datas, "", "    ")
	//Error handling
	if err != nil {
		log.Println(err)
	}

	//Write it to the file
	err = ioutil.WriteFile("./people/all.json", dataBytes, 0644)
	//Error handling
	if err != nil {
		log.Println(err)
	}

	result := `{"status":200, "message":"Noice"}`
	var finalResult map[string]interface{}
	json.Unmarshal([]byte(result), &finalResult)

	json.NewEncoder(response).Encode(finalResult)
}

func CreateAppointment(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	var appointment Appointment
	err = json.Unmarshal(body, &appointment)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(appointment)

	file, _ := json.MarshalIndent(appointment, "", "    ")

	_ = ioutil.WriteFile("./appointments/"+fmt.Sprint(appointment.ID)+".json", file, 0644)

	allfile, err := ioutil.ReadFile("./appointments/all.json")
	if err != nil {
		log.Println(err)
	}

	datas := []Appointment{}

	json.Unmarshal(allfile, &datas)

	//Define what we want to add
	newStruct := &Appointment{
		ID:        appointment.ID,
		FirstName: appointment.FirstName,
		LastName:  appointment.LastName,
		Date:      appointment.Date,
		Time:      appointment.Time,
		Reason:    appointment.Reason,
		DiscordID: appointment.DiscordID,
	}

	datas = append(datas, *newStruct)

	//JSON-lize the data defined above
	dataBytes, err := json.MarshalIndent(datas, "", "    ")
	//Error handling
	if err != nil {
		log.Println(err)
	}

	//Write it to the file
	err = ioutil.WriteFile("./appointments/all.json", dataBytes, 0644)
	//Error handling
	if err != nil {
		log.Println(err)
	}

	result := `{"status":200, "message":"Noice"}`
	var finalResult map[string]interface{}
	json.Unmarshal([]byte(result), &finalResult)

	json.NewEncoder(response).Encode(finalResult)
}

func GetApp(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	if _, err := os.Stat("./appointments/" + id + ".json"); err == nil {
		json_file, err := os.Open("./appointments/" + id + ".json")
		if err != nil {
			log.Fatal(err)
		}
		defer json_file.Close()

		byteValue, _ := ioutil.ReadAll(json_file)

		var appointment Appointment

		json.Unmarshal(byteValue, &appointment)
		fmt.Println(appointment)
		json.NewEncoder(response).Encode(appointment)

	} else if os.IsNotExist(err) {
		result := `{"status":404, "message":"File does not exist. Person has no appointment."}`
		var finalResult map[string]interface{}
		json.Unmarshal([]byte(result), &finalResult)

		json.NewEncoder(response).Encode(finalResult)
	}
}

func main() {
	fmt.Println("Starting the api....")
	route := mux.NewRouter()
	router := cors.Default().Handler(route)

	route.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	route.HandleFunc("/people", GetAllPeople).Methods("GET")
	route.HandleFunc("/createApp", CreateAppointment).Methods("POST")
	route.HandleFunc("/getApp/{id}", GetApp).Methods("GET")
	//route.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	//route.HandleFunc("/rmperson/{id}", DeletePersonEndpoint).Methods("DELETE")
	//route.HandleFunc("/changeperson/{id}/{type}/{value}", UpdatePersonEndpoint).Methods("PATCH")
	route.NotFoundHandler = http.HandlerFunc(whereyoulooking)
	http.ListenAndServe(":12345", router)
}
