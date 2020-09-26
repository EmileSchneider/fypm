package session

import "fypm.com/domain/entity"

type repository interface {
	Create(s *Session) (entity.ID, error) 
	Delete(id entity.ID) error
}

type Manager interface {
	repository 
}
