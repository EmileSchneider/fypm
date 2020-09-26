package client

import "fypm.com/domain/entity"

type Reader interface{
	Get(id entity.ID) (*Client, error)
	List() ([]*Client, error)
}

type Writer interface {
	Create(c *Client) (entity.ID, error)
	Update(c *Client) error
	Delete(id entity.ID) error
}

type repository interface {
	Reader
	Writer
}

type Manager interface {
	repository
}
