package main

import (
	"github.com/walbety/go-funds-htmx/cmd/rest"
)

type Count struct {
	Count int
}

func main() {

	// config.Initialize()

	rest.StartRest()

}
