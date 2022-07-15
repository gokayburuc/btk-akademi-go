package conditionals

import "fmt"

func Workshop() {
	// uc adet degisken tanimlayin
	// ekrana en büyük değişkeni yazdirin

	// var degisken1 int = 25
	// var degisken2 int = 9
	// var degisken3 int = 11

	// var (
	// 	degisken4 int = 13
	// 	degisken5 int = 18
	// )

	// uc degisken tanimlama duz yazim
	var deg, deg1, deg2 int = 11, 25, 9

	var max = 0

	if deg > max {
		max = deg
		if deg1 > max {
			max = deg1
		} else if deg2 > max {
			max = deg2
		}
	}

	fmt.Printf("En buyuk deger : %v", max)
}
