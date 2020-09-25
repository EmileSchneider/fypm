package testdbsetup

import (
	"fypm.com/domain/entity"
	"fypm.com/domain/entity/user"
	"fypm.com/domain/entity/broker"
)

type setupusecase struct {
	userManager user.Manager
	brokerManager broker.Manager
}

func NewUseCase(u user.Manager, b broker.Manager) *setupusecase {
	return &setupusecase{
		userManager: u,
		brokerManager: b,
	}
}

func (u *setupusecase) Setup() {
	for i := 1; i < 10; i++ {
		user := user.User{
			ID: entity.NewID(),
			Email: "mail" + string(i) + "@mail.com",
			Password: "unqiue" + string(i*i) + "password",
		}
		broker := broker.Broker{
			ID: entity.NewID(),
			User: user.ID,
			FirstName: "Hans" + string(i),
			LastName: "Wurst" + string(i*i),
		}
		u.userManager.Create(&user)
		u.brokerManager.Create(&broker)
	}
}


