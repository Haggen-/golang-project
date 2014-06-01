package routing

import (
	"log"
	"net/http"
)

const PACKAGE_NAME string = "Routing"

type WebModule interface {
	ModuleInformation
	SubModules() []SubModule
	ModuleRoutes
}

type ModuleRoutes interface {
	Get(r *http.Request, res http.ResponseWriter) error
	Post(r *http.Request, res http.ResponseWriter) error
	Put(r *http.Request, res http.ResponseWriter) error
	Delete(r *http.Request, res http.ResponseWriter) error
}

type SubModule interface {
	ModuleInformation
	ModuleRoutes
}

type ModuleInformation interface {
	Version() string
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
		webServerModules[module.Version()] = module		
	} else {
		log.Println("Attempted to add nil value to webServerModule list, nil value discarded.")
	}
}