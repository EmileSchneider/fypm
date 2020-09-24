package user

import "fypm.com/domain/entity"

type User struct {
	ID entity.ID
	Email string
	Password string
}
