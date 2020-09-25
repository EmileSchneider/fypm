package auth

import (
	"fypm.com/domain"
	"fypm.com/domain/entity"
	"fypm.com/domain/entity/user"
)

type usecase struct {
	uManager user.Manager
}

func NewUseCase(u user.Manager) *usecase {
	return &usecase{
		uManager: u,
	}
}

func (ucs *usecase) Authenticate(name, password string) error {
	return nil
}
