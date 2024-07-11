package service

import (
	"time-tracker/internal/model"
	"time-tracker/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) GetUsers(filters map[string]string, limit int, cursor int) ([]model.User, error) {
	return s.repo.GetUsers(filters, limit, cursor)
}

func (s *UsersService) GetUserById(userId int) (model.User, error) {
	return s.repo.GetUserById(userId)
}

func (s *UsersService) CreateUser(user model.User) (userId int, err error) {
	return s.repo.CreateUser(user)
}

func (s *UsersService) DeleteUser(userId int) (err error) {
	return s.repo.DeleteUser(userId)
}

func (s *UsersService) UpdateUser(userId int, userDataToUpdate model.UpdateUserInput) (model.User, error) {

	err := s.repo.UpdateUser(userId, userDataToUpdate)

	if err != nil {
		return model.User{}, err
	}

	user, err := s.repo.GetUserById(userId)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
