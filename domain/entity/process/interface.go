package process

import "fypm.com/domain/entity"

type Reader interface {
	Get(id entity.ID) (*Process, error)
	List() ([]*Process, error)
}

type Writer interface {
	Create(p *Process) (entity.ID, error)
	Update(p *Process) error
	Delete(id entity.ID) error
}

type repository interface {
	Reader
	Writer
}

type Manager interface {
	repository
}
