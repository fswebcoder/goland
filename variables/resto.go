package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var estado bool
var sueldo float32
var fecha time.Time

func RestoVariables() {
	fmt.Println(time.Now())
}

func ConvietoAtexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
