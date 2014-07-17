package users

import (
	"encoding/json"
	"log"
	"fmt"
	"net/http"
	router "github.com/haggen-/nexus-events/routing"
	)

type User struct {
	Firstname string
	Surname string
	Ssn string
}

type UserModule struct {
	router.WebModule
}

func (t UserModule) GetPath() string {
	return "/user/"
}

func (t UserModule) GetVersion() string {
	return "1.0"
}

func (t UserModule) HandleRequest(res http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t.get(res, r)
	case "POST":
		t.post(res, r)
	case "PUT":
		t.put(res, r)
	case "DELETE":
		t.delete(res, r)
	default:
		log.Printf("Unknown method %v request made towards userModule")
	}
}

func (t UserModule) post(res http.ResponseWriter, r *http.Request) {
	log.Println("POST request made!")
} 

func (t UserModule) get(res http.ResponseWriter, r *http.Request) {
	log.Println("GET request made!")
	usr := User { "John", "Boehner", "66666-5566" }
	usr2 := User { "Jake", "Bacon", "6121231"}

	result := [2]User{usr, usr2}
	jsonUsr, ok := json.Marshal(result)
	if(ok != nil) {
		log.Println("Unexpected error occurred.2")
	}
    res.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(res, string(jsonUsr))
}

func (t UserModule) put(res http.ResponseWriter, r *http.Request) {
	log.Println("PUT request made!")
}

func (t UserModule) delete(res http.ResponseWriter, r *http.Request) {
	log.Println("DELETE request made!")
}

func init() {
	module := UserModule { }
	log.Println("Package 'users' initalized and added to Module List.");
	router.RegisterModule(module)
}