package funciones

import (
	"fmt"
)

func Calculos() {
	/*
		FUNCIÓN ANONIMA
	*/
	calculo := func(numero1 int, numero2 int) int {
		return numero1 * numero2
	}
	fmt.Println(calculo)
}
