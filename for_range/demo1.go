package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "Bursa", "Kirklareli"}

	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	// range yazimi
	for _, sehir := range sehirler {
		fmt.Println(sehir)
	}

	//index range
	for i, sehir := range sehirler {
		fmt.Println(i, "-", sehir)
	}
}
