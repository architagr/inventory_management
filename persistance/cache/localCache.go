package cache

import "fmt"

type LocalCache struct {
}

func (cache *LocalCache) getValue(keyName string) interface{} {
	return fmt.Sprintf("%s_value", keyName)
}
