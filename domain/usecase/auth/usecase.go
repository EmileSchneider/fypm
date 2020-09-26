package auth

import (
	"fypm.com/domain"
	"fypm.com/domain/entity"
	"fypm.com/domain/entity/user"
	"fypm.com/domain/entity/session"
)

type usecase struct {
	uManager user.Manager
	sManager session.Manager
}

func NewUseCase(u user.Manager, s session.Manager) *usecase {
	return &usecase{
		uManager: u,
		sManager: s,
	}
}

func (ucs *usecase) Authenticate(email, password string) (string, error) {
	user := uManager.GetByEmail(email)
	if user.Password != password {
		return nil, errors.New("Wrong Password")
	}
	return nil
}
