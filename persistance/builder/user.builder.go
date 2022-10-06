package builder

import (
	"inventory_management/model"
	"inventory_management/persistance"
	"inventory_management/persistance/contracts"
	"inventory_management/persistance/dynamodb"
	"inventory_management/persistance/mongodb"
)

type UserPersistanceFactory struct{}

func (*UserPersistanceFactory) InitUsersPersistance(dbType, userId int) contracts.IUser {
	dbConn := persistance.GetUserConnection(userId)
	if dbType == persistance.DynamoDbPersistance {
		return newUserDynamodbBuilder(dbConn)
	} else if dbType == persistance.MongoDbPersistance {
		return newUserMongodbBuilder(dbConn)
	}
	return nil
}

func newUserDynamodbBuilder(dbConn model.DbConnection) contracts.IUser {
	return dynamodb.UserConstructor(dbConn)
}
func newUserMongodbBuilder(dbConn model.DbConnection) contracts.IUser {
	return mongodb.UserConstructor(dbConn)
}
