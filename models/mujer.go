package models

type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar()    { m.Respiracion = true }
func (m *Mujer) Comer()       { m.Comiendo = true }
func (m *Mujer) Pensar()      { m.Pensando = true }
func (m *Mujer) Sexo() string { return "Mujer" }
