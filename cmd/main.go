package main

import (
	"fmt"

	math_lib "github.com/ImKairat/golang-lab/internal/math_lib"
)

func main() {
	rand_num, err := math_lib.GenRandNumberRange(990, 1000)
	if err != nil {
		fmt.Println("Occured error:", err)
		return
	}
	fmt.Printf("Generated random number: %d\n", rand_num)
}
