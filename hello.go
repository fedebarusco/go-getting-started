package main

import (
	"fmt"

	"rsc.io/quote"

	"go-getting-started/values"

	"go-getting-started/variables"

	"go-getting-started/constants"
)

func main() {
	fmt.Println(quote.Go())
	values.PrintValues()
	variables.PrintVariables()
	constants.PrintConstants()
}
