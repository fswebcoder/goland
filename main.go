package main

import (
	"github.com/fswebcoder/goland/middleware"
)

func main() {

	// if os := runtime.GOOS; os == "Linux." || os == "OS X." {
	// 	fmt.Println("esto no es windows")
	// } else {
	// 	fmt.Println("esto es windows")
	// 	fmt.Println(os)
	// }

	// switch os := runtime.GOOS; os {
	// case "Linux.":
	// 	fmt.Println("Esto es linux")

	// case "windows":
	// 	fmt.Println("Esto es windows desde el case")
	// default:
	// 	fmt.Printf("%s \n ", os)

	// }

	// fmt.Println(ejercicios.Conversor("1000"))

	// teclado.IngresoNumeros()

	// iteraciones.Iterar()

	// Pedro := new(models.Mujer)
	// Pedro.Edad = 20
	// Pedro.Altura = 1.84
	// Pedro.Peso = 84
	// Pedro.Vivo = true

	// ejerinterfaces.HumanosRespirando(Pedro)
	// canal1 := make(chan bool)
	// go goroutines.MiNombreLentooo("Fabio", canal1)
	// defer func() {
	// 	<-canal1
	// }()

	// fmt.Println("Estoy aqui")

	// webserver.MiWebServer()

	middleware.MiMiddleware()
}
