package variables

import "fmt"

func PrintVariables() {

	var a = "initial"
	fmt.Println("Variable a =", a)

	var b, c int = 1, 2
	fmt.Println("Variables b and c =", b, c)

	var d = true
	fmt.Println("Variable d =", d)

	var e int
	fmt.Println("Variable e =", e)

	f := "apple"
	fmt.Println("Variable f =", f)

}
