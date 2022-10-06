package services

import (
	"inventory_management/model"
	"inventory_management/persistance/builder"
)

func InitUserService(persistanceBuilder builder.UserPersistanceFactory) IUserService {
	return &UserService{
		userPersistanceBuilder: persistanceBuilder,
	}
}

var _ IUserService = new(UserService)

type IUserService interface {
	GetAllUsers(dbType, loggedInUserId int) []model.User
}
type UserService struct {
	userPersistanceBuilder builder.UserPersistanceFactory
}

func (svc *UserService) GetAllUsers(dbType, loggedInUserId int) []model.User {
	return svc.userPersistanceBuilder.InitUsersPersistance(dbType, loggedInUserId).GetUsers()
}
