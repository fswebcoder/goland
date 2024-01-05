package ejerinterfaces

import (
	"fmt"

	"github.com/fswebcoder/goland/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.Sexo())

}
