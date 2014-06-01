package main

import (
		"log"
		"net/http"
		_ "github.com/haggen-/nexus-events/routing"
		_ "github.com/haggen-/nexus-events/modules/users"
		)

type HttpSettings struct {
	portNumber string
}

type Module struct {
	name string
	initalize func() error
}

func main() {
	log.Println("Hello world!!");
	settings := HttpSettings { ":8080" }
	log.Println("What " + settings.portNumber)
	err := startWebServer(settings);
	if(err != nil) {
		log.Fatal(err)
	}
	log.Println("Webserver started, listening to port " + settings.portNumber)
}

func startWebServer(settings HttpSettings) error {
	return http.ListenAndServe(settings.portNumber, nil)
}

