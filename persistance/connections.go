package persistance

import "inventory_management/model"

func GetUserConnection(userId int) model.DbConnection {
	// TODO: get dbconnection from cache else get from a static db
	return userId
}
