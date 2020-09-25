package broker

import (

	"fypm.com/domain/entity"
)

type manager struct {
	repo repository
}

func NewManager(repo repository) *manager {
	return &manager{
		repo: repo,
	}
}

func (m *manager) Create(b *Broker) (entity.ID, error) {
	b.ID = entity.NewID()
	return m.repo.Create(b)
}

func (m *manager) Get(id entity.ID) (*Broker, error) {
	return m.repo.Get(id)
}



func (m *manager) Update(b *Broker) error {
	return m.repo.Update(b)
}

func (m *manager) Delete(id entity.ID) error {
	_ , err := m.repo.Get(id)
	if err != nil {
		return err
	}
	return m.repo.Delete(id)
}

func (s *manager) List() ([]*Broker, error) {
	return s.repo.List()
}

