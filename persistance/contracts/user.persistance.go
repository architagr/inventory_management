package contracts

import (
	"inventory_management/model"
)

type IUser interface {
	GetUsers() []model.User
}
