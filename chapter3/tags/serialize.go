package tags

import (
	"fmt"
	"reflect"
)

// SerializeStructStrings converts a struct
// to our custom serialization format
// it honors serialize struct tags for string types
func SerializeStructStrings(s interface{}) (string, error) {
	result := ""

	//reflect the interface into a type
	r := reflect.TypeOf(s)
	fmt.Println(r)
	value := reflect.ValueOf(s)
	fmt.Println(value)

	// if a pointer to a struct is passed
	// handle it appropriately
	// TODO: check wheither this works
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		value = value.Elem()
	}

	// loop over all the fields
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		// struct tag found
		// key := field.Name // example was bugged
		// TODO: add other types to serialize too
		// key := ""

		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// ignore "-" otherwise the whole value becomes the serialize key
			if serialize == "-" {
				continue
			}
			key := serialize
			fmt.Println("serialize: ", serialize, "key: ", key, "ok", ok, "type: ", value.Field(i).Kind())

			switch value.Field(i).Kind() {
			// TODO:  Only supports string in this example
			// The example was bugged
			case reflect.String:
				result += key + ":" + value.Field(i).String() + ";"
			case reflect.Int:

				output := fmt.Sprint(value.Field(i).Int())
				// fmt.Println(key, field.Type, value.Field(i))
				result += key + ":" + output + ";"
				// result += key + ":" + value.Field(i).Int() + ";"
				// result += key + ":" + value.Field(i) + ";"
				// fmt.Println(key)
			default:
				continue
			}
		}
	}
	return result, nil
}
