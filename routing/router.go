package routing

import (
	"fmt"
	"log"
	"net/http"
	"errors"
)

const PACKAGE_NAME string = "Routing"

type WebModule interface {
	GetVersion() string
	GetPath() string

	Get(r *http.Request, res http.ResponseWriter) error
	Post(r *http.Request, res http.ResponseWriter) error
	Put(r *http.Request, res http.ResponseWriter) error
	Delete(r *http.Request, res http.ResponseWriter) error
}

var webServerModules = make(map[string]WebModule)

func init() {
	log.Println("Package '" + PACKAGE_NAME + "' initalized.")
}

func NumberOfModules() int {
	return len(webServerModules)
}

func RegisterModule(module WebModule) {
	if(module != nil) {
		webServerModules[module.GetPath()] = module		
	} else {
		log.Println("Attempted to add nil value to webServerModule list, nil value discarded.")
	}
}

func HandleRequest(path string, r *http.Request, res http.ResponseWriter) error {
	val, ok := webServerModules[path]
	if(!ok) {
		message := fmt.Sprintf("Could not handle request for path %v, no appropriate webmodule found.", path)
		log.Println(message)
		return errors.New(message)
	}

	log.Println("Received %v request to handler.")	
	switch r.Method {
	case "GET":
		return val.Get(r, res)
	case "PUT":
		return val.Put(r, res)
	case "POST":
		return val.Post(r, res)
	case "DELETE":
		return val.Delete(r, res)
	default:
		return errors.New(fmt.Sprintf("Could not handle request using method %v", r.Method))
	}
}