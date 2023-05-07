package slices

import "fmt"

func PrintValues() {

	// Slices are typed only by the elements they contain;
	// To create an empty slice with non-zero length, use the builtin make
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator
	// The following example gets a slice of the elements s[2], s[3], s[4]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Gets a slice up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// Gets a slice up from (but including) s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// Declare and initialize a variable for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices can be composed into multi-dimensional data structures
	// The length of the inner slices can vary
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
