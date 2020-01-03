package log

import (
	"bytes"
	"fmt"
	"log"
)

// Log uses the setup logger
func Log() {
	// write logs to bytes.Buffer

	buf := bytes.Buffer{}

	// second argument is the prefix
	// last argument is about options you combine with logical OR

	logger := log.New(&buf, "Logger: ", log.Lshortfile|log.Ldate)

	logger.Println("Test")

	logger.SetPrefix("New Logger: ")
	logger.Printf("you can also add args(%v) and use Fataln to log and crash", true)
	fmt.Println(buf.String())
}
