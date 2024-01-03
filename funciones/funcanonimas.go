package funciones

import (
	"fmt"
)

func calculos() {
	/*
		FUNCIÓN ANONIMA
	*/
	calculo := func(numero1 int, numero2 int) int {
		return numero1 * numero2
	}
	fmt.Println(calculo)
}

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {
	tabladel2 := 2
	Mitabla := tabla(tabladel2)
	for i := 1; i <= 10; i++ {
		fmt.Println(Mitabla())
	}
}
