package errorhandling

import "fmt"

func Dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)

}

func Demo3() {
	var sayi interface{} = "Engin"
	Dogrula(sayi)

	var sayi2 interface{} = 25
	Dogrula(sayi2)

}
