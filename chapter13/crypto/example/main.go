package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/jaakidup/go-cookbook/chapter13/crypto"
)

func main() {
	http.HandleFunc("/guess", crypto.GuessHandler)
	fmt.Println("Running on localhost:8080")
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}
