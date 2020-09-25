package broker

import "fypm.com/domain/entity"

type Reader interface {
	Get(id entity.ID) (*Broker, error)
	List() ([]*Broker, error)
}

type Writer interface {
	Create(b *Broker) (entity.ID, error)
	Update(b *Broker) error
	Delete(id entity.ID) error
}

type repository interface {
	Reader
	Writer
}

type Manager interface {
	repository
}
