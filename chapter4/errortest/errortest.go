package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	CheckSystem()
}

// CheckSystem ...
func CheckSystem() {
	es := NewErrorsSystem()
	defer es.CheckErrors()
	es.AddCheck(sample(false, "first function"))
	es.Must(sample(false, "Must Do this"))
	es.AddCheck(sample(true, "third function"))
}

func deferer() {
	var err error
	defer func() {
		fmt.Println("Running error check")
		// fmt.Println(err)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	sample(false, "first function")
	err = sample(false, "second function")
	err = sample(true, "third function")

	// fmt.Println(RetError())
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func sample(shouldFail bool, message string) error {

	if shouldFail {
		return errors.New(message)
	}
	return nil
}

// RetError ...
func RetError() error {
	err := sample(true, "RetError test")
	return errors.Wrap(err, "This only does something if err != nil")
}
