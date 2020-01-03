package main

import (
	"fmt"

	"github.com/jaakidup/go-cookbook/chapter3/currency"
)

func main() {

	userInput := "15.93"
	pennies, err := currency.ConvertDolllarsToPennies(userInput)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User input converted to %d pennies\n", pennies)
	// adding 15 cents
	pennies += 15
	dollars := currency.ConvertPenniesToDollarString(pennies)
	fmt.Printf("Added 15 cents, new values is %s dollars\n",
		dollars)

}
