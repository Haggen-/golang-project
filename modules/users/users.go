package users

import (
	"log"
	router "github.com/haggen-/nexus-events/routing"
	)

type HttpSettings struct {
	portNumber string
}

type Module struct {
	name string
	initalize func() error
}

func init() {
	log.Println("Package 'users' initalized and added to Module List.");
	router.RegisterModule(nil)
}