package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
)

type Person struct {
    ID string `json:"id,omitempty"`
    FirstName string `json:"firstname,omitempty"`
    LastName string `json:"lastname,omitempty"`
    Address *Address `json:"address,omitempty"`
}
type Address struct {
    City string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}
var people []Person


func GetPeopleEndpoint(w http.ResponseWriter, request *http.Request){
    log.Println("Personas:",len(people))
    json.NewEncoder(w).Encode(people)
}
func GetPersonEndpoint(w http.ResponseWriter, request *http.Request){
    params := mux.Vars(request)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}
func SetPersonEndpoint(w http.ResponseWriter, request *http.Request){
    //params := mux.Vars(request)
    var person Person
    _ = json.NewDecoder(request.Body).Decode(&person)

    // convertimos el id de la ultima persona a entero, le sumamos 1 y a string de nuevo
    lastID, err := strconv.Atoi(people[len(people)-1].ID)
    if err == nil {
        lastID++
        person.ID = strconv.Itoa(lastID)
    }
    //person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}
func DelPersonEndpoint(w http.ResponseWriter, request *http.Request){
    params := mux.Vars(request)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index + 1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(people)
}

func main(){
    log.Println("Starting server...")
    router := mux.NewRouter()

    // adding example data
    people = append(people, Person{ID: "1", FirstName:"Ryan", LastName:"Reynolds", Address: &Address{City:"Los Angeles", State:"California"}})
    people = append(people, Person{ID: "2", FirstName:"Maria", LastName:"Alonso"})

    // endpoints
    router.HandleFunc("/people"     , GetPeopleEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
    router.HandleFunc("/people"     , SetPersonEndpoint).Methods("POST")
    router.HandleFunc("/people/{id}", DelPersonEndpoint).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":3000", router))
}