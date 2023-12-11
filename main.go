package seolist

import (
	"fmt"
	"seolist/keywords"
)

func main() {

	_, err := keywords.FromFile("test.csv", keywords.Processor)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
