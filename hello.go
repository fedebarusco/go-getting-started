package main

import (
	"fmt"
	"go-getting-started/arrays"
	"go-getting-started/constants"
	"go-getting-started/forloop"
	"go-getting-started/ifelse"
	"go-getting-started/maps"
	"go-getting-started/ranges"
	"go-getting-started/slices"
	"go-getting-started/switch_statement"
	"go-getting-started/values"
	"go-getting-started/variables"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	values.PrintValues()
	variables.PrintVariables()
	constants.PrintConstants()
	forloop.PrintValues()
	ifelse.PrintValues()
	switch_statement.PrintValues()
	arrays.PrintValues()
	slices.PrintValues()
	maps.PrintValues()
	ranges.PrintValues()
}
