package main

import "fmt"

// CheckType prints the interface type
func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("default")
	}
}

func main() {
	CheckType("Jaaki")
	CheckType(3)
	CheckType(true)

	type Person struct {
		Name string
	}

	var i interface{}
	i = Person{Name: "jaaki"}
	if value, ok := i.(bool); ok {
		fmt.Println("value", value)
	}

	if val, ok := i.(string); ok {
		fmt.Println("Val is", val)
	}

	if val, ok := i.(Person); ok {
		fmt.Println("Val is type Person", val)
	}

}

// func InterfaceTypeCheck(i interface{}) {
// 	switch s.(type) {
// 	case string:
// 		return
// 	case int:
// 		fmt.Println("int")
// 	default:
// 		fmt.Println("default")
// 	}
// }
