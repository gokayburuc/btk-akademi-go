package deferstatement

import (
"fmt"
)

func OddEven(sayi int32) string {
	defer DeferFunc()
	
	if sayi%2 == 0 {
		return "Cift sayidir"
	}
	if sayi%2 != 0 {
		return "Tek sayidir"
	}

	return "Belli degil"
}


func DeferFunc(){
	fmt.Println("Defer Func Calisti.İslem Sona Erdi.")
}

func Test(){
	yanit := OddEven(23)
	fmt.Println(yanit)

}