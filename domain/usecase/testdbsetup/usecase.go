package testdbsetup

import (
	"fmt"
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
			Email: "mail" + "22" + "@mail.com",
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

func (u *setupusecase) Test(){
	brokerList, err := u.brokerManager.List()
	if err != nil {
		fmt.Println(err)
	}
	for _ , ur  := range brokerList {
		uu := ur
		uu.FirstName = "HANSIGERWURSTER"
		err = u.brokerManager.Update(uu)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Broker FirstName before: %s, after: %s \n", ur.FirstName, uu.FirstName)
	}
}


