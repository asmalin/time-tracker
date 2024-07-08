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
