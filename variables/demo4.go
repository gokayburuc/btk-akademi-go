package variables

import (
	"fmt"
)

func Demo4() {
	x := 10
	y := 25.3

	fmt.Println(float64(x) + y)
	fmt.Println(x + int(y))

	fmt.Printf("%v %T", x, x)
	fmt.Printf("%v %T", y, y)

	var p int8 = 17
	var z int16 = 23

	fmt.Print("\n")
	fmt.Println(int16(p) + z)

	var g float32
	var s float32 = 11
	fmt.Println(s + g)

	veri := 106

	// formatted text yazdirir
	fmt.Printf("%v %T \n", veri, veri)

	//ascii kodu karsiligi olan j degerini getirir
	fmt.Println(string(veri))

}
