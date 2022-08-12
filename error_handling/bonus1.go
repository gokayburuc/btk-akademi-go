package errorhandling

import (
	"fmt"
	"log"
	"os"
)

func OpenFile() {
	file, err := os.Open("demo1.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(file)
	}
}
