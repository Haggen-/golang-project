package routing

import (
	"log"
	"net/http"
)

const PACKAGE_NAME string = "Routing"

type WebModule interface {
	GetVersion() string
	GetPath() string

	HandleRequest(http.ResponseWriter, *http.Request)

	get(res http.ResponseWriter, r *http.Request)
	post(res http.ResponseWriter, r *http.Request)
	put(res http.ResponseWriter, r *http.Request)
	delete(res http.ResponseWriter, r *http.Request)
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

func InitPaths() error {
	for path, module := range webServerModules {
		http.HandleFunc(path, module.HandleRequest)
	}
	return nil
}