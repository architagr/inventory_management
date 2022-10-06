package main

import (
	"fmt"
	"inventory_management/persistance"
	"inventory_management/persistance/builder"
	"inventory_management/services"
)

func main() {
	userPerBuilder := builder.UserPersistanceFactory{}
	userService := services.InitUserService(userPerBuilder)

	fmt.Printf("users in dynamoDB %+v\n", userService.GetAllUsers(persistance.DynamoDbPersistance, 1))
	fmt.Printf("users in mongodb %+v\n", userService.GetAllUsers(persistance.MongoDbPersistance, 1))

}
