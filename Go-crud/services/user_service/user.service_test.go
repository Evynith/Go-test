package user_service

import (
	"testing"

	m "go-crud/models"
)

func TestCreate(t *testing.T) {
	user := m.User{
		Name:  "Evy3",
		Email: "evynith@gmail.com",
	}
	err := Create(user) //user_service esta implicito porque esta en el mismo paquete

	if err != nil {
		t.Error("La prueba de persistencia de datos ha fallado")
		t.Fail()
	} else {
		t.Log("Se realizó con éxito")
	}
}

/*
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
*/
func TestReadOne(t *testing.T) {
	user, err := ReadOne("63063b9aae872a4c4b147dde")

	if err != nil {
		t.Error("hubo un error")
		t.Fail()
	} else {
		t.Log("exito!", user)
	}
}

func TestUpdate(t *testing.T) {
	user := m.User{
		Name:  "evy 2",
		Email: "email.falso@123",
	}

	err := Update(user, "000000000000000000000000")

	if err != nil {
		t.Error("hubo un error")
		t.Fail()
	} else {
		t.Log("exito!")
	}
}

func TestDelete(t *testing.T) {
	err := Delete("000000000000000000000000")

	if err != nil {
		t.Error("hubo un error")
		t.Fail()
	} else {
		t.Log("exito!")
	}
}
