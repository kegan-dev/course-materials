package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

)

type Response struct{
	Assignments []Assignment `json:"assignments"`
}

type ClassResponse struct{
	Classes []Class `json:"classes"`
}

type Assignment struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"desc"`
	Points int `json:"points"`
}

type Class struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Instructor string `json:"instr"`
	Description string `json:"desc"`
	Room string `json:"room"`
}

var Assignments []Assignment
var Classes []Class

func InitAssignments(){
	var assignmnet Assignment
	assignmnet.Id = "Mike1A"
	assignmnet.Title = "Lab 4 "
	assignmnet.Description = "Some lab this guy made yesteday?"
	assignmnet.Points = 20
	Assignments = append(Assignments, assignmnet)
}

func InitClasses(){
	var class Class
	class.Id = "BHG"
	class.Title = "Black Hat Go"
	class.Description = "Golang hacking."
	class.Instructor = "Mike Borowczak"
	class.Room = "CR500"
	Classes = append(Classes, class)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}


func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	for _, assignment := range Assignments {
		if assignment.Id == params["id"]{
			json.NewEncoder(w).Encode(assignment)
			break
		}
	}
	//TODO : Provide a response if there is no such assignment
	//w.Write(jsonResponse)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
			if assignment.Id == params["id"]{
				Assignments = append(Assignments[:index], Assignments[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s PUT end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	params := mux.Vars(r)		
	found := false
	for ix, a := range Assignments {
		if a.Id == params["id"] {
			Assignments[ix].Id = r.FormValue("id")
			Assignments[ix].Title = r.FormValue("title")
			Assignments[ix].Description = r.FormValue("desc")
			Assignments[ix].Points, _ = strconv.Atoi(r.FormValue("points"))
			found = true
			break;
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	var response Response
	response.Assignments = Assignments
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	var assignmnet Assignment
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("id") != ""){
		assignmnet.Id =  r.FormValue("id")
		assignmnet.Title =  r.FormValue("title")
		assignmnet.Description =  r.FormValue("desc")
		assignmnet.Points, _ =  strconv.Atoi(r.FormValue("points"))
		Assignments = append(Assignments, assignmnet)
		w.WriteHeader(http.StatusCreated)
		return
	}
	w.WriteHeader(http.StatusNotFound)

}

// Classes Endpoints
func GetClasses(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response ClassResponse

	response.Classes = Classes

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s POST end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	var class Class
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("id") != ""){
		class.Id =  r.FormValue("id")
		class.Title =  r.FormValue("title")
		class.Description =  r.FormValue("desc")
		class.Instructor =  r.FormValue("instr")
		class.Room = r.FormValue("room")
		Classes = append(Classes, class)
		w.WriteHeader(http.StatusCreated)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, class := range Classes {
			if class.Id == params["id"]{
				Classes = append(Classes[:index], Classes[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}
