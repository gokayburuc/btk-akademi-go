package deferstatement

import "fmt"

func A() {
	fmt.Println("A Fonskiyonu çalıştı.")
}

func C() {
	fmt.Println("C Fonksiyonu çalıştı.")
}

func D() {
	fmt.Println("D Fonksiyonu çalıştı.")
}

func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("B Fonskiyonu çalıştı.")
}
