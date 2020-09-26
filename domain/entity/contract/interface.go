package contract

import "fypm.com/domain/entity"

type Reader interface {
	Get(id entity.ID) (*Contract, error)
	List() ([]*Contract, error)
}

type Writer interface {
	Create(c *Contract) (entity.ID, error)
	Update(c *Contract) error
	Delete(id entity.ID) error
}

type repository interface {
	Reader
	Writer
}

type Manager interface {
	repository
}
