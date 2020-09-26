package process

import "fypm.com/domain/entity"

type Process struct {
	ID, Client, Contract entity.ID
	Status string
}
