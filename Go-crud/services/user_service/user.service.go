package user_service

import (
	m "go-crud/models"
	userRepository "go-crud/repositories/user_repository"
)

func Create(user m.User) error {
	err := userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func Read(name string, date1 string, date2 string) (m.Users, error) {

	users, err := userRepository.Read(name, date1, date2)

	if err != nil {
		return nil, err
	}
	return users, nil
}

func ReadOne(userId string) (m.User, error) {

	user, err := userRepository.ReadOne(userId)
	if err != nil {
		return m.User{}, err
	}
	return user, nil
}

func Update(user m.User, userId string) error {
	err := userRepository.Update(user, userId)
	if err != nil {
		return err
	}
	return nil
}

func Delete(userId string) error {
	err := userRepository.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
