package main

import "github.com/jaakidup/go-cookbook/chapter1/filedirs"

func main() {

	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	if err := filedirs.CapitalizerExample(); err != nil {
		panic(err)
	}
}
