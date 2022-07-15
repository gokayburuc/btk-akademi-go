package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 10000
	var cekilecekmeblag float64 = 9000

	if cekilecekmeblag > hesap {
		fmt.Println("Yetersiz Bakiye!")

	} else if cekilecekmeblag == hesap {
		fmt.Println("Banka bosaltildi!")
	} else if cekilecekmeblag <= hesap {
		fmt.Println("Al parani Arap!")
	} else {
		fmt.Println("Sana para Yok Arap!")
	}

	fmt.Println("Uygulama Sonlandi!")

}
