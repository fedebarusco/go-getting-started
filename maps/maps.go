package maps

import "fmt"

func PrintValues() {

	// To create an empty map, use the builtin make(map[key-type]value-type)
	m := make(map[string]int)

	// Set key/value pairs
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key doesn't exists the zero value of the value type is returned
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	// delete removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	// If the the value itself isn't needed, it can be ignored with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// A new map can be declared and initialized in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
