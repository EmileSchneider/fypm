package session

import "fypm.com/domain/entity"

type manager struct {
	repo repository
}

func NewManager(repo repository) *manager {
	return &manager{
		repo: repo, 
	}
}

func (m *manager) Create(s *Session) (entity.ID, error){
	session := Session{ID: entity.NewID()}
	err := m.repo.Create(session)
	if err != nil {
		return nil, err
	}
	return session.ID, nil
}

func (m *manager) Delete(s *Session) error {
	return m.repo.Delete(s.ID)
}
