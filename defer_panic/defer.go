package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Abriendo Archivo")
	defer fmt.Println("Esto se ejecuta al final")
	fmt.Println("Leyendo Archivo")
}

func VemosUnPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ups, me mori %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Encontre un error")
	}
}
