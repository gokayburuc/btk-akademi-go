package loops

import "fmt"

// kullanıcıdan sayi girmesini iste
// giriline sayi asal mi kontrol et
// 23  asal sayidir. Yalnizca kendisine ve 1 e bolunur

func AsalSayi() {
	sayi := 0
	fmt.Println("Bir sayi giriniz: \n")
	fmt.Scanln(&sayi) // buraya bir pointer ekledik

	asalAnahtar := true
	for i := 2; i < sayi; i++ {

		if sayi%i == 0 {
			asalAnahtar = false
			fmt.Printf("Sayi : %v Mod :%v \n", i, sayi%i)
		}

	}

	if !asalAnahtar {
		fmt.Println("Asal Degil")
	}

	if asalAnahtar == true {
		fmt.Println("Asal")

	}

}
