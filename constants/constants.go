package constants

import (
	"fmt"
	"math"
)

const s string = "constant"

func PrintConstants() {

	fmt.Println("Constant s =", s)

	const n = 500000000

	const d = 3e20 / n

	fmt.Println("Constant d =", d)

	fmt.Println("Int64(d) =", int64(d))

	fmt.Println("Sin(d) =", math.Sin(n))

}
