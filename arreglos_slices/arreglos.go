package arreglos_slices

import (
	"fmt"
)

var tabla [10]int //array de 10 posiciones
var matriz [2][2]int

func MuestraArreglos() {
	// tabla[0] = 1
	// tabla[5] = 15
	matriz[0][0] = 1
	matriz[0][1] = 2
	matriz[1][0] = 3
	matriz[1][1] = 4
	// fmt.Println(tabla)
	fmt.Println(matriz)
}
