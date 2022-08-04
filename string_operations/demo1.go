package string_operations

import (
	"fmt"
	"strings"
)

func Demo1() {

	name := "Gokay BURUC"
	fmt.Println(strings.Replace(name, "a", "z", -1))
	fmt.Println(strings.Compare(name, "Gökay BÜRÜÇ"))
	fmt.Println(strings.Count(name, "U"))

	// Join func
	nameArray := []string{"gok", "ay", "buruc"}
	fmt.Println(strings.Join(nameArray, "-"))

	// Has prefix
	fmt.Println(strings.HasPrefix(name, "Gok"))

	// Has suffix
	fmt.Println(strings.HasSuffix(name, "RUC"))

	// Index
	fmt.Println(strings.Index(name, "R"))

	// ToUpper
	fmt.Println(strings.ToUpper(name))

	// ToTitle
	fmt.Println(strings.ToTitle(name))

	// ToLower
	fmt.Println(strings.ToLower(name))

	// Repeat
	fmt.Println(strings.Repeat("=", 5))

	// Split
	fmt.Println(strings.Split(name, " "))

	// Contains
	fmt.Println(strings.Contains(name, "y"))

	// Last Index
	fmt.Println(strings.LastIndex(name, "U"))

	// Trim
	fmt.Println(strings.Trim(" x asdasd lhdfgdfgdiş     ", " "))

	// Fields
	fmt.Println(strings.Fields(name))

}
