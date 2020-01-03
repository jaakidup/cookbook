package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SearchString shoes a number of methods
// for searching a string
func SearchString() {
	s := "this is a test"
	// returns true because s contains
	// the word this
	fmt.Println(strings.Contains(s, "this"))

	// returns true because s contains the letter a
	// would also math if it contained b or c
	fmt.Println(strings.ContainsAny(s, "abc"))

	// return tru because s starts with this
	fmt.Println(strings.HasPrefix(s, "this"))

	// return true becuase s ends with test
	fmt.Println(strings.HasSuffix(s, "test"))
}

// ModifyString modifies a string in a number of ways
func ModifyString() {
	s := "simple string"
	// prints [simple string]
	fmt.Println(strings.Split(s, " "))
	// prints "Simple String"
	fmt.Println(strings.Title(s))
	// prints "simple string"; all trailing and
	// leading white space is removed
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

// StringsReader demostrates how to create an io.Reader interface
// quickly with a string
func StringsReader() {
	s := "simple string"
	r := strings.NewReader(s)
	io.Copy(os.Stdout, r)
}
