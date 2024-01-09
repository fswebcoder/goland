package middleware

import (
	"fmt"
)

func sumar(a int, b int) int {
	return a + b
}

func restar(a int, b int) int {
	return a - b
}

func multiplicar(a int, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Esto es un middleware")
	resultado := operaciones(restar)(10, 20)
	print(resultado)
}

func operaciones(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
