package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fswebcoder/goland/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.Tabla()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Tabla()
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_RDONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteStrinng" + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeerArchivos() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Hubo un error leyendo archivos" + err.Error())
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(">" + registro)
	}
}
