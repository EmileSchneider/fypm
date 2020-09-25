package broker

import "fypm.com/domain/entity"

type Broker struct {
	ID entity.ID
	User entity.ID
	FirstName, LastName string
	Processes []*entity.ID
}
