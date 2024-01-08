package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Moive struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json: title`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var moives []Moive

func main() {
	r := mux.NewRouter()

	moives = append(moives, Moive{ID: "1", Isbn: "438227", Title: "Tanasha", Director: &Director{Firstname: "john", Lastname: "Doe"}})
	moives = append(moives, Moive{ID: "2", Isbn: "125435", Title: "Ye Jawani Hai Diwani", Director: &Director{Firstname: "Aryan", Lastname: "Mukargi"}})

	r.HandleFunc("/moives", getMoives).Methods("GET")
	r.HandleFunc("/moives/{id}", getMoive).Methods("GET")
	r.HandleFunc("/moives", createMoive).Methods("POST")
	r.HandleFunc("/moives/{id}", updateMoive).Methods("PUT")
	r.HandleFunc("/moives/{id}", deleteMoive).Methods("DELETE")

	fmt.Print("web is serving on the port: 8084")
	log.Fatal(http.ListenAndServe(":8084", r))

}

/*
jsondata , err := json.Marshal(moives);

		if err != nil{
			log.Fatal(err);
		}
	   fmt.Print("moives are ",jsondata);
*/
func getMoives(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(moives)

}

func deleteMoive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range moives {
		if item.ID == params[`id`] {
			moives = append(moives[:index], moives[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(moives)

}

func getMoive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	found := false
	for _, item := range moives {
		if params[`id`] == item.ID {
			found = true
			break
		}
	}

	if !found {
		errMess := "enter the valid ID"
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errMess)
		return
	}
	json.NewEncoder(w).Encode(moives)
}

func createMoive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var moive Moive
	_ = json.NewDecoder(r.Body).Decode(&moive)
	moive.ID = strconv.Itoa(rand.Intn(100000000))
	moives = append(moives, moive)

	json.NewEncoder(w).Encode(moive)
}
func updateMoive(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//param
	params := mux.Vars(r)
	//json over the moives, range
	//delete the moive i.d that you ahve send in body of postman
	for index, item := range moives {
		if item.ID == params["id"] {
			moives = append(moives[:index], moives[index+1:]...)
			var moive Moive
			_ = json.NewDecoder(r.Body).Decode(&moive)
			moive.ID = params["id"]
			moives = append(moives, moive)
			json.NewEncoder(w).Encode(moive)
			return
		}

	}

}
