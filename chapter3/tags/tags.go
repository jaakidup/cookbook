package tags

import "fmt"

// Person is a struct that stores a persons
// name, city, state, and a misc attribute
type Person struct {
	Name  string `serialize:"name"`
	City  string `serialize:"city"`
	State string `serialize:"state"`
	Misc  string `serialize:"-"`
	Year  int    `serialize:"year"`
	Cool  bool   `serialize:"cool"`
}

// EmptyStruct demonstrates serialize
// and deserialize for an Empty struct
// with tags
func EmptyStruct() error {
	p := Person{}
	res, err := SerializeStructStrings(&p)
	if err != nil {
		return err
	}

	fmt.Printf("Empty struct: %#v\n", p)
	fmt.Println("Serialize Results:", res)
	newP := Person{}
	if err := DeSerializeStructStrings(res, &newP); err != nil {
		return err
	}
	fmt.Printf("Deserialize results: %#v\n", newP)
	return nil
}

// FullStruct demonstrates serialize
// and deserialize for an Full struct
// with tags
func FullStruct() error {
	p := Person{
		Name:  "Aaron",
		City:  "Seattle",
		State: "WA",
		Misc:  "some fact",
		Year:  2017,
		Cool:  true,
	}
	res, err := SerializeStructStrings(&p)
	if err != nil {
		return err
	}
	fmt.Printf("Full struct: %#v\n", p)
	fmt.Println("Serialize Results:", res)
	// newP := Person{}
	// if err := DeSerializeStructStrings(res, &newP); err != nil {
	// 	return err
	// }
	// fmt.Printf("Deserialize results: %#v\n", newP)
	return nil
}
