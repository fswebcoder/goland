package models

type Hombre struct {
	Edad        int
	Altura      float32
	Peso        float32
	Respiracion bool
	Pensando    bool
	Comiendo    bool
	Vivo        bool
}

func (h *Hombre) Respirar()    { h.Respiracion = true }
func (h *Hombre) Comer()       { h.Comiendo = true }
func (h *Hombre) Pensar()      { h.Pensando = true }
func (h *Hombre) Sexo() string { return "Hombre" }
