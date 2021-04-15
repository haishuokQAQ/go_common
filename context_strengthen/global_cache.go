package context_strengthen

import (
    `context`
    `sync`
)

var globalCacheKey = &struct {}{}

func createGlobalCacheContext(parent context.Context) context.Context {
    return context.WithValue(parent, globalCacheKey, new(sync.Map))
}

func GetCacheMap(ctx context.Context) *sync.Map {
    mapInterface := ctx.Value(globalCacheKey)
    if mapInterface == nil {
        return nil
    }
    return mapInterface.(*sync.Map)
}