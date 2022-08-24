package user_service

import (
	"testing"
	"time"

	m "go-crud/models"
)

func TestCreate(t *testing.T) {
	user := m.User{
		Name:      "Evy",
		Email:     "evynith@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := Create(user) //user_service esta implicito porque esta en el mismo paquete

	if err != nil {
		t.Error("La prueba de persistencia de datos ha fallado")
		t.Fail()
	} else {
		t.Log("Se realizó con éxito")
	}
}

func TestRead(t *testing.T) {
	users, err := Read()

	if err != nil {
		t.Error("hubo un error")
		t.Fail()
	}
	if len(users) == 0 {
		t.Error("No hay elementos para mostrar")
		t.Fail()
	} else {
		t.Log("exito!")
	}
}

func TestUpdate(t *testing.T) {
	user := m.User{
		Name:  "evy 2",
		Email: "email.falso@123",
	}

	err := Update(user, "0")

	if err != nil {
		t.Error("hubo un error")
		t.Fail()
	} else {
		t.Log("exito!")
	}
}

func TestDelete(t *testing.T) {
	err := Delete("0")

	if err != nil {
		t.Error("hubo un error")
		t.Fail()
	} else {
		t.Log("exito!")
	}
}