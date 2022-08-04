package constants

import (
	"fmt"
	"math"
)

func Demo1() {

	const pi = 3.14
	yaricap := 12.0

	fmt.Println(yaricap * yaricap * pi)
	fmt.Println(math.Pow(float64(yaricap), 2) * pi)

}
