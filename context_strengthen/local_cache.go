package context_strengthen

import (
    `context`
    `sync`
)

var localCacheKey = &struct {}{}

func CreateLocalCacheContext(parent context.Context) context.Context {
    return context.WithValue(parent, localCacheKey, new(sync.Map))
}

func LocalCacheMap(ctx context.Context) *sync.Map {
    mapInterface := ctx.Value(localCacheKey)
    if mapInterface == nil {
        return nil
    }
    return mapInterface.(*sync.Map)
}
