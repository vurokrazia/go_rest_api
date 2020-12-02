package main
import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
	"go_api_rest/connect"
	"go_api_rest/structures"
	// "fmt"
	)

func main()  {
	connect.InitializeDatabase()	
	defer connect.CloseConnection()
	r := mux.NewRouter()
	r.HandleFunc("/user/", GetUsers).Methods("GET")
	r.HandleFunc("/user/", NewUser).Methods("POST")
	r.HandleFunc("/user/{id}", UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	log.Println("The server found in 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func GetUser(w http.ResponseWriter, r* http.Request){
	vars := mux.Vars(r)
	user_id := vars["id"]
	user := connect.GetUser(user_id)
	status := "success"
	var message string
	if user.Id == 0 {
		status = "error"
		message = "User not found"
	}
	response := structures.Response{ status, user, message }
	DisplayResponse(w, response )
}

func GetUsers(w http.ResponseWriter, r* http.Request)  {
	user := connect.GetUsers()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func NewUser(w http.ResponseWriter, r* http.Request)  {
	user := GetUserRequest(r)
	response := structures.Response{ "success", connect.CreateUser(user), "" }
	DisplayResponse(w, response )
}

func UpdateUser(w http.ResponseWriter, r* http.Request)  {
	vars := mux.Vars(r)
	user_id := vars["id"]

	user := GetUserRequest(r)
	response := structures.Response{ "success", connect.UpdateUser(user_id, user), "" }
	DisplayResponse(w, response )
}

func DeleteUser(w http.ResponseWriter, r* http.Request)  {
	vars := mux.Vars(r)
	user_id := vars["id"]

	connect.DeleteUser(user_id)
	DisplayResponse(w, structures.Response{ "success", structures.User{} , "" })
}

func GetUserRequest(r* http.Request) structures.User  {
	var user = structures.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func DisplayResponse(w http.ResponseWriter, response structures.Response)  {
	log.Println(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}