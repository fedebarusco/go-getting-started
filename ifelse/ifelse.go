package ifelse

import "fmt"

func PrintValues() {

	// basic example
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// else statement is not mandatory
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// statement can precede contitions;
	// the defined variable(s) is/are available(s) in the current and all subsequet branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "is single digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
