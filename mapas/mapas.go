package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string, 5)
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	paises["Colombia"] = "Bogota"
	paises["Peru"] = "Lima"
	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
		"America":      26,
		"Milan":        18,
		"Bayern":       15,
		"Junior":       27,
	}
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s tiene un puntaje de: %d \n", equipo, puntaje)
	}

	delete(campeonato, "Real Madrid") // Elimina un elemento del mapa
}
