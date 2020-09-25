package user

import "fypm.com/domain/entity"
/*
type Reader interface {
	Get(id entity.ID) (*User, error)
	Search(query string) ([]*User, error)
	List() ([]*User, error)
}

type Writer interface {
	Create(u *User) (entity.ID, error)
	Update(u *User) error
	Delete(id entity.ID) error
}

type repository interface {
	Reader
	Writer
}
*/

type repository interface {
	Get(id entity.ID) (*User, error)
	Create(u *User) (entity.ID, error)
	GetByMail(mail string) (*User, error)
}

type Manager interface {
	repository
}
