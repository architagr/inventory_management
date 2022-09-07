package contracts

type ICommonPersistance interface {
	GetConnection(userId int) interface{}
}

type CommonPersistance struct {
}

func (per *CommonPersistance) GetConnection(userId int) interface{} {
	return 1
}
