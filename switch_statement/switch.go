package switch_statement

import (
	"fmt"
	"time"
)

func PrintValues() {

	// Basic switch
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Use commas to separate multiple expressions in the same statement
	// Use optional default case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch without an expression is an alternative way to express if/else logic
	switch t := time.Now().Hour(); {
	case t < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// swtich can also compares types insted of values
	// the example discover the type of an interface value
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
