package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("error_handling/demo1.txt")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Dosya bulunamadi")
	} else {
		fmt.Println(f.Name())
	}
}
