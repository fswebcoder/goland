package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string, 5)
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	paises["Colombia"] = "Bogota"
	paises["Peru"] = "Lima"
	fmt.Println(paises)
}
