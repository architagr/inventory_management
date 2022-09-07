package mongodb

import (
	"inventory_management/model"
	"inventory_management/persistance/contracts"
)

type User struct {
	contracts.CommonPersistance
	connection interface{}
}

func UserConstructor(userId int) User {
	commonPer := contracts.CommonPersistance{}
	return User{
		CommonPersistance: commonPer,
		connection:        commonPer.GetConnection(userId),
	}
}

var _ contracts.IUser = new(User)

func (user User) GetUsers() []model.User {
	return nil
}
