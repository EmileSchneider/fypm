package client

import "fypm.com/domain/entity"

type Client struct {
	ID entity.ID
	FirstName, LastName, Address, Country string
}
