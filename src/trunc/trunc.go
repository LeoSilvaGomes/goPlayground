package trunc

import (
	"fmt"
	"math"
)

func Trunc() {
	var x float64

	fmt.Println("Write a float number:")
	fmt.Scan(&x)

	t := math.Trunc(x)

	fmt.Println("Truncated value: ", t)
}
