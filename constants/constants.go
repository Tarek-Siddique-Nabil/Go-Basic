package constants

import (
	"fmt"
	"math"
)

const s string = "constant"

func constants() {
	fmt.Println(s)

	var n int

	fmt.Print("Type N value: ")
	fmt.Scan(&n)

	// d is now a variable because its value depends on the runtime value of n
	d := 3e20 / float64(n)

	fmt.Println(d)

	// Use int64 instead of int32 to avoid overflow
	fmt.Println(int64(d))

	// Calculate the sine of n
	fmt.Println(math.Sin(float64(n)))
}
