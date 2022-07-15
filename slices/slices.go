package slices

import "fmt"

func Slice1() {
	isimler := make([]string, 4)
	fmt.Println(isimler)

	isimler[0] = "Gokay"
	isimler[1] = "Esin"
	isimler[2] = "Guler"
	isimler[3] = "Sukru"

	fmt.Println(isimler)

	isimler = append(isimler, "Berk")

	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
