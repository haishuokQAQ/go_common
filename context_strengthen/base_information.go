package context_strengthen

import (
    `context`
    `github.com/haishuokQAQ/go_common/util`
)

var basementContextKey = &struct {}{}

type GetBaseInfoMethod func()map[interface{}]interface{}

func createBaseInfoContext(getMethod GetBaseInfoMethod) context.Context{
    result := getMethod()
    unmodifiableMap := util.NewUnmodifiableMap(result)
    resultCtx := context.WithValue(context.Background(), basementContextKey, unmodifiableMap)
    return resultCtx
}

func GetBasementValue(ctx context.Context, key interface{}) (interface{}, bool) {
    unmodifiableMapInter := ctx.Value(basementContextKey)
    if unmodifiableMapInter == nil {
        return nil, false
    }
    return unmodifiableMapInter.(*util.UnmodifiableMap).Load(key)
}

func ForceGetBasementValue(ctx context.Context, key interface{})interface{}{
    value, _ := GetBasementValue(ctx, key)
    return value
}


