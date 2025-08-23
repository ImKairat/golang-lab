package datatypes

import (
	"fmt"
	"reflect"
)

func SimpleDataTypes() {
	numeric_datatypes()
	runes_and_bytes()
}

func numeric_datatypes() {
	a, b := 1, 2
	c := 3
	var z int32
	z = 12

	fmt.Printf("%d, %d, %d, %d\n", a, b, c, z)
}

func runes_and_bytes() {
	name := "❌Gopher"
	fmt.Println(reflect.TypeOf(name))

	name2 := []rune(name)
	fmt.Printf("%c \n", name2[0])
	fmt.Println(reflect.TypeOf(name2).Kind())
}
