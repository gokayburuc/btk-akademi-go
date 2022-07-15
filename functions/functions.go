package functions

import "fmt"

func Topla(sayi1 int, sayi2 int) int {
	var toplam = sayi1 + sayi2
	fmt.Printf("%v + %v ", sayi1, sayi2)
	fmt.Println("Sonuç: ", toplam)
	return toplam

}

func SelamVer() {
	fmt.Println("Merhabalar")
}

func KullaniciSelamla(username string) {
	fmt.Println("Merhaba, ", username)
}
