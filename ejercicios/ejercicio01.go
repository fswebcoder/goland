package ejercicios

import (
	"strconv"
)

func Conversor(valor string) (int, string) {
	num, _ := strconv.Atoi(valor)
	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}
