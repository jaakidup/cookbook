package main

import (
	"fmt"
	"github.com/jaakidup/go-cookbook/chapter7/controllers"
	"net/http"
)

func main() {
	storage := controllers.MemStorage{}
	c := controllers.NewController(&storage)
	http.HandleFunc("/get", c.GetValue(false))
	http.HandleFunc("/get/default", c.GetValue(true))
	http.HandleFunc("/set", c.SetValue)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe("127.0.0.1:3333", nil)
	panic(err)
}
