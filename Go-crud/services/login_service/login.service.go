package login_service

import (
	auth "go-crud/repositories/auth_repository"
)

func LoginUser(email string, password string) bool {
	user, err := auth.SearchUser(email)

	if err != nil {
		return false
	}
	return user.Password == password
}
