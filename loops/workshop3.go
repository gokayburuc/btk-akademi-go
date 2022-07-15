package loops

import "fmt"

// 220 ve 284 arkadaş sayılardır
// 10 ve 65 arkadaş sayi değildir

func ArkadasSayi() {
	sayi1 := 220
	sayi2 := 284
	toplam1 := 0
	toplam2 := 0

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			toplam1 = toplam1 + i

		}
	}

	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			toplam2 = toplam2 + i
		}
	}

	fmt.Printf("Sayi1 : %v \,n", sayi1)
	fmt.Printf("Sayi2 : %v \n", sayi2)
	fmt.Printf("Toplam1 : %v \n", toplam1)
	fmt.Printf("Toplam1 : %v \n", toplam2)

	if toplam1 == sayi2 && toplam2 == sayi1 {
		fmt.Printf("%v ve %v arkadaş sayilardir.", sayi1, sayi2)
	} else {
		fmt.Printf("%v ve %v arkadaş sayılar değildir.", sayi1, sayi2)
	}

}
