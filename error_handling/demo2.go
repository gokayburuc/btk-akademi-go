package errorhandling

import (
	"fmt"
	"os"
)

func CatchError() {
	f, err := os.Open("demo2.txt")

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya : ", pErr.Path)

		} else {
			fmt.Println(" Dosya Bulunamadi")
		}

		fmt.Println(f.Name())
	}
}
