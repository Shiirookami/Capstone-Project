package service

import (
	"GOLANG/entity"
	"context"
	"errors"
)

type LoginUseCase interface {
	Login(ctx context.Context, email string, password string) (*entity.User, error)
}

type LoginRepository interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type loginService struct {
	repository LoginRepository
}

func NewLoginService(repository LoginRepository) *loginService {
	return &loginService{
		repository: repository,
	}
}

// func untuk melakikan pengecekan apakah semua data nya sama mulai dari email, password
func (s *loginService) Login(ctx context.Context, email string, password string) (*entity.User, error) {
	user, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	//untuk pengecakan apakah email  ada di database?
	if user == nil {
		return nil, errors.New("user with that email not found")
	}

	//untuk pengecekan apakah password nya ada atau gaa di databse?
	if user.Password != password {
		return nil, errors.New("incorrect login credentials")
	}

	//ketika email dan passwerd sama maka akan mengembalikan nil
	return user, nil

}
