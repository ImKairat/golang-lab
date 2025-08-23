package datatypes

import "fmt"


func Pointers() {
	x, y, z := 23, 3.14, "Hi!"
	
	var (
		px *int = &x
		py *float64 = &y
		pz *string = &z
	)

	print_pointer(ToPrint[int]{px, "Until change"})
	*px = 46
	print_pointer(ToPrint[int]{px, "After change"})

	print_pointer(ToPrint[float64]{p: py})
	print_pointer(ToPrint[string]{p: pz})

	print_pointer(ToPrint[string]{})
}


type ToPrint[T any] struct{
	p *T
	message string
}

func print_pointer[T any](tp ToPrint[T]) {
	if tp.p == nil {
		fmt.Println("[WARNING] Pointer is nil!")
		return
	}

	if tp.message == "" {
		fmt.Printf("Value: %v\nPointer: %p\n\n", *tp.p, tp.p)
	} else {
		fmt.Printf("Message: %v\nValue: %v\nPointer: %p\n\n", tp.message, *tp.p, tp.p)
	}
}
