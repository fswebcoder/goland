package models

import (
	"time"
)

type User struct {
	Id               int
	Nombre           string
	Fecha_nacimiento time.Time
	Create_at        time.Time
	Status           bool
}

func (this *User) AddUser(id int, nombre string, fecha_nacimiento time.Time, Create_at time.Time, Status bool) {
	this.Id = id
	this.Nombre = nombre
	this.Fecha_nacimiento = fecha_nacimiento
	this.Create_at = Create_at
	this.Status = Status
}
