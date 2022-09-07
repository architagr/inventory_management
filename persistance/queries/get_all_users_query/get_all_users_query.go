package get_all_users_query

import (
	"inventory_management/model"
	"inventory_management/persistance/contracts"
	"inventory_management/persistance/dynamodb"
	"inventory_management/persistance/mongodb"
	"inventory_management/persistance/queries"
)

type IGetAllUsersQuery interface {
	Query() []model.User
}

type GetAllUsersQuery[T contracts.IUser] struct {
	userPersistance T
}

func (q GetAllUsersQuery[T]) Query() []model.User {
	return q.userPersistance.GetUsers()
}

func Constructor(dbType, userId int) IGetAllUsersQuery {
	if dbType == queries.DynamoDbPersistance {
		return NewGetAllUsersQueryDynamodb(userId)
	} else if dbType == queries.MongoDbPersistance {
		return NewGetAllUsersQueryMongodb(userId)
	}
	return nil
}
func NewGetAllUsersQueryDynamodb(userId int) IGetAllUsersQuery {
	return GetAllUsersQuery[dynamodb.User]{
		userPersistance: dynamodb.UserConstructor(userId),
	}
}
func NewGetAllUsersQueryMongodb(userId int) IGetAllUsersQuery {
	return GetAllUsersQuery[mongodb.User]{
		userPersistance: mongodb.UserConstructor(userId),
	}
}

var _ IGetAllUsersQuery = new(GetAllUsersQuery[dynamodb.User])
