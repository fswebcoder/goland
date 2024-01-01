package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

// Ejercicio 02
// Crear un archivo GO en el paquete 'ejercicios' llamado 'ejercici002.go'
// Crear una función Pública, que pida por teclado un número, valide si hay
// error o no (y si hay error que vuelva a pedirlo).
// En base al número recibido crear la tabla numérica del mismo. De 1 a 10 y la
// muestre por pantalla
// En Main, llamar a dicha función
// Grabar, ejecutar y luego subir todo a Github
func Tabla() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese un número entero")

	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			Tabla()
		}
		fmt.Println("***************")
		for i := 1; i <= 10; i++ {
			numeroComoCadena := strconv.Itoa(numero)
			iComoCadena := strconv.Itoa(i)
			multiplicacionComoCadena := strconv.Itoa(numero * i)
			resultado := numeroComoCadena + " x " + iComoCadena + " = " + multiplicacionComoCadena
			fmt.Println(resultado)
		}
	}
	Tabla()

}
