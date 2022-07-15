package slices

import "fmt"

func Slices2() {
	sehirler := []string{"Ankara", "Bursa", "İstanbul"}
	fmt.Println(sehirler)

	//sehirler kopya
	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)

	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)

	sehirler = append(sehirler, "Adana")
	fmt.Println(sehirlerKopya)
	fmt.Println(sehirler)

	// dizinin 3 ve 4 arası
	fmt.Println(sehirler[3:4])

	//bastan 2 ye dek
	fmt.Println(sehirler[:2])

	//
	fmt.Println(sehirler[2:])
}
