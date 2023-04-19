package main

import (
	"fmt"

	"rsc.io/quote"

	"go-getting-started/values"

	"go-getting-started/variables"

	"go-getting-started/constants"

	"go-getting-started/forloop"

	"go-getting-started/ifelse"
)

func main() {
	fmt.Println(quote.Go())
	values.PrintValues()
	variables.PrintVariables()
	constants.PrintConstants()
	forloop.PrintValues()
	ifelse.PrintValues()
}
