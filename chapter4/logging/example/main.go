package main

import (
	"fmt"

	log "github.com/jaakidup/go-cookbook/chapter4/logging"
)

func main() {
	fmt.Println("Basic loggin and modification of logger")

	log.Log()
	fmt.Println("logging 'handled' errors:")

	log.FinalDestination()
}
