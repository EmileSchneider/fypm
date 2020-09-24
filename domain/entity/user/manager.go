package user

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

func (m *manager) Create(u *User) (entity.ID, error) {
	u.ID = entity.NewID()
	return m.repo.Create(u)
}

func (m *manager) Get(id entity.ID) (*User, error) {
	return m.repo.Get(id)
}
/*
func (m *manager) Update(u *User) error {
	return m.repo.Update(u)
}

func (m *manager) Delete(id entity.ID) error {
	u1, err := m.repo.Get(id)
	if err != nil {
		return err
	}
	return m.repo.Delete(id)
}
func (s *manager) Search(query string) ([]*User, error) {
	return s.repo.Search(strings.ToLower(query))
}

func (s *manager) List() ([]*User, error) {
	return s.repo.List()
}
*/
