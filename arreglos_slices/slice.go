package arreglos_slices

import (
	"fmt"
)

var tablas [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //array de 10 posiciones

func MuestraSlice() {
	porcuatro := tablas[3:]
	fmt.Println(porcuatro)

	porcuatro1 := tablas[:5]
	fmt.Println(porcuatro1)

	porcuatro2 := tablas[3:6]
	fmt.Println(porcuatro2)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
