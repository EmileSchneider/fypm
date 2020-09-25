package auth

import(
	"fypm.com/domain/entity/user"
)

type UseCase interface {
	Authenticate(name, password string) error
	Logout(u *User) error
}
