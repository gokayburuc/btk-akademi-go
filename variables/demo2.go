package variables

import "fmt"

func Demo2() {

	var (
		name            string = "Gokay"
		surname         string = "BURUC"
		age             int    = 35
		bachelor_degree bool   = true
	)

	fmt.Printf("%s %s %v %v ", name, surname, age, bachelor_degree)
}
