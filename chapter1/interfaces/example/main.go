package main

import (
	"bytes"
	"fmt"
	"github.com/jaakidup/go-cookbook/chapter1/interfaces"
)

func main() {

	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Println("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("out bytes buffer = ", out.String())
	fmt.Println("stdout on PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
	println()
}
