package mongodb

import (
	"inventory_management/model"
	"inventory_management/persistance/contracts"
)

type User struct {
	contracts.CommonPersistance
	connection model.DbConnection
}

func UserConstructor(dbConn model.DbConnection) User {
	commonPer := contracts.CommonPersistance{}
	return User{
		CommonPersistance: commonPer,
		connection:        dbConn,
	}
}

var _ contracts.IUser = new(User)

func (user User) GetUsers() []model.User {
	return []model.User{
		{
			UserId:    "MongoDb UserId",
			FirstName: "MongoDb FirstName",
		},
	}
}
