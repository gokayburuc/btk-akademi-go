package arrays

import "fmt"

func Sehirler() {

	var sehirler [7]string

	sehirler[0] = "Ankara"
	sehirler[1] = "İstanbul"
	sehirler[2] = "Bursa"
	sehirler[3] = "İzmir"
	sehirler[4] = "Edirne"
	sehirler[5] = "Tekirdag"
	sehirler[6] = "Kırklareli"

	fmt.Println(sehirler)

	for i := 0; i < 7; i++ {
		fmt.Println(i, ": ", sehirler[i])
	}
}
