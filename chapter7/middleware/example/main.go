package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jaakidup/go-cookbook/chapter7/middleware"
)

func main() {
	// We apply from the bottom up
	h := middleware.ApplyMiddleware(
		middleware.Handler,
		middleware.Logger(log.New(os.Stdout, "", 0)),
		middleware.SetID(100),
	)
	http.HandleFunc("/", h)
	fmt.Println("Listening on port :3333")
	log.Fatalln(http.ListenAndServe(":3333", nil))
}
