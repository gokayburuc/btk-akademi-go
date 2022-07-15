package loops

import "fmt"

// tahmin sayisina bagli ifadeler
// 1-3 tahmin : Aferin Arap!
// 4-6 tahmin : Zordu ama Bildin Arap!
// 6 > tahmin : Amma cahil adamsın Arap!

func SayiBul() {
	tutulanSayi := 17
	tahminEdilenSayi := 0
	tahminSayisi := 0

	fmt.Println("Bir Sayi Tahmin Edin : ")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi += 1

	for tutulanSayi != tahminEdilenSayi {
		if tutulanSayi < tahminEdilenSayi {
			fmt.Println("Yuh deve! Ucma!")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi += 1
		} else if tutulanSayi > tahminEdilenSayi {
			fmt.Println("Cik daha! Cik!")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi += 1
		} else if tutulanSayi == tahminEdilenSayi {
			fmt.Println("Aferin bildin len!")
		}
	}

	basariDurumu := ""

	if tahminSayisi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Bravo Arap!"
	}
	if tahminSayisi >= 4 && tahminSayisi <= 6 {
		basariDurumu = "Çok zorlama Motoru Bozarsın Arap!"
	}
	if tahminSayisi > 6 {
		basariDurumu = "Amma Cahil adamsın Arap!"
	}

	fmt.Printf("Tahmin Sayisi : %v\n", tahminSayisi)
	fmt.Print(basariDurumu)

}
