package main

import (
	"github.com/jaakidup/go-cookbook/chapter3/tags"
)

func main() {
	// if err := tags.EmptyStruct(); err != nil {
	// 	panic(err)
	// }
	// fmt.Println()
	if err := tags.FullStruct(); err != nil {
		panic(err)
	}
}
