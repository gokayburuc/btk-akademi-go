package conditionals

import "fmt"

func Demo3() {
	var para float64 = 4000.00
	var borc float64 = 5000.00
	if para > 3000 {
		if para <= borc {
			fmt.Println("Battin Arap!")
		} else {
			fmt.Println("Bu ay da ölmedin Arap!")
		}
		fmt.Println("Tebrikler!")
	}
}
