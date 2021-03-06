package contract

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

func (m *manager) Create(c *Contract) (entity.ID, error) {
	c.ID = entity.NewID()
	return m.repo.Create(c)
}


func (m *manager) Get(id entity.ID) (*Contract, error) {
	return m.repo.Get(id)
}

func (m *manager) Update(c *Contract) error {
	return m.repo.Update(c)
}

func (m *manager) Delete(id entity.ID) error {
	_ , err := m.repo.Get(id)
	if err != nil {
		return err
	}
	return m.repo.Delete(id)
}

func (s *manager) List() ([]*Contract, error) {
	return s.repo.List()
}


