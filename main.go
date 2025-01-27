package main

import (
	"fmt"
	"log"

	"github.com/sarveshraj/testcodecov/mypackage"
)

func main() {
	result, err := mypackage.ConvIntToStr(3)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(result)
}
