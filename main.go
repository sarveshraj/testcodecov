package main

import (
	"fmt"

	"github.com/sarveshraj/testcodecov/mypackage"
)

func main() {
	result, err := mypackage.ConvIntToStr(3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
