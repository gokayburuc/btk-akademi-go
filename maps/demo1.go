package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)

	sozluk["bursa"] = "16"
	sozluk["izmir"] = "35"
	sozluk["Kırklareli"] = "39"

	fmt.Println(sozluk["bursa"])
	fmt.Println("Eleman Sayisi", len(sozluk))

	fmt.Println(sozluk)
	delete(sozluk, "bursa")
	fmt.Println("Eleman Sayisi", len(sozluk))
	fmt.Println(sozluk)

	deger := sozluk["izmir"]
	// sozlukte deger aramak
	fmt.Println(deger)

	deger, mevcut := sozluk["izmir"]

	fmt.Println(deger)
	fmt.Println("Listede var mi : ", mevcut)

	//
	deger2, mevcut2 := sozluk["istanbul"]
	fmt.Println(deger2)
	fmt.Println("Listede var mi : ", mevcut2)
}
