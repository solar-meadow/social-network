package service

import "github.com/expose443/social-network/internal/repository"

type AuthServiceI interface {
	Register()
}

type authService struct {
	repo repository.AuthRepositoryI
}

func (a *authService) Register() {
	a.repo.Register()
}
