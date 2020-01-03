package main

import "fmt"

import "github.com/pkg/errors"

import "log"

// NewErrorsSystem returns a pointer to a new ErrorsSystem
// with an initialized errors map
func NewErrorsSystem() *ErrorsSystem {
	es := &ErrorsSystem{}
	es.counter = 0
	es.errors = make(map[int]error)
	return es
}

// ErrorsSystem ...
type ErrorsSystem struct {
	counter int
	errors  map[int]error
}

// AddCheck to add an error for later checking
func (es *ErrorsSystem) AddCheck(err error) {
	es.counter++
	es.errors[es.counter] = err
}

// CheckErrors to check all the add errors
func (es *ErrorsSystem) CheckErrors() {
	if es.counter == 0 {
		fmt.Println("No errors to check")
	}
	for i, err := range es.errors {
		if err != nil {
			fmt.Println("Error: ", i, err.Error())
			fmt.Println(errors.Cause(err).Error())
			fmt.Println(errors.WithStack(err))
		}
	}
}

// Must execute this function, else panic
func (es *ErrorsSystem) Must(err error) {
	if err != nil {
		log.Fatalln("Failed on : ", err.Error())
	}
}
