package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 10000
	var cekilecekmeblag float64 = 9000

	if cekilecekmeblag > hesap {
		fmt.Println("Yetersiz Bakiye")
	}

	if cekilecekmeblag <= hesap {
		fmt.Println("Bakiyeniz :", hesap)
		//fmt.Println("Cekilecek Tutar : ", cekilecekmeblag)
		fmt.Println("Cekilecek Tutar : " + fmt.Sprintf("%v", cekilecekmeblag))
		bakiye := hesap - cekilecekmeblag
		fmt.Printf("Kalan Meblag : %v \n", bakiye)
	}

	fmt.Println("Uygulama Sonlandi!")

}
