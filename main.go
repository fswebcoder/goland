package main

import (
	"fmt"

	"github.com/fswebcoder/goland/variables"
)

func main() {
	variables.RestoVariables()
	estado, texto := variables.ConvietoAtexto(15)
	fmt.Println(estado)
	fmt.Println(texto)
}
