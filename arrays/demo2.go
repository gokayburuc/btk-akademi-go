package arrays

import "fmt"

func ArrayDemo() {
	sayilar := [5]int{20, 30, 50, 10, 2}
	fmt.Println(sayilar)

	for i := 0; i < len(sayilar); i++ {
		fmt.Println(sayilar[i])
	}

	max := 0
	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > max {
			max = sayilar[i]
		}
	}

	fmt.Printf("MAX : %v \n", max)
}
