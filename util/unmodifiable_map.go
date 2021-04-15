package util

import (
    `sync`
)

type UnmodifiableMap struct {
    *sync.Map
}

func (um *UnmodifiableMap) Store(key, value interface{}) {
    return
}

func (um *UnmodifiableMap) Delete(key interface{}) {
    return
}

func NewUnmodifiableMap(data map[interface{}]interface{})*UnmodifiableMap{
    innerMap := new(sync.Map)
    for key, value := range data {
        innerMap.Store(key, value)
    }
    return &UnmodifiableMap{innerMap}
}

