package session

import "fypm.com/domain/entity"

type Session struct {
	ID entity.ID
	User entity.ID
	Secret string
}
