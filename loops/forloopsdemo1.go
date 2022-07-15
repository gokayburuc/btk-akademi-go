package loops

import "fmt"

func Demo() {

	var metin string = " Selamın Aleyküm! "

	i := 1
	for i <= 5 {
		fmt.Println(metin)
		i += 1
	}

	fmt.Print("Sonlandi!")
}
