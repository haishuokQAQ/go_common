package context_strengthen

import (
    `context`
    `github.com/opentracing/opentracing-go`
)

var globalGetBaseMethod GetBaseInfoMethod

var defaultGetBaseMethod GetBaseInfoMethod = func() map[interface{}]interface{} {
    return nil
}

func SetGlobalGetBaseMethod(method GetBaseInfoMethod) {
    globalGetBaseMethod = method
}

func InitBasementContext() context.Context {
    method := globalGetBaseMethod
    if globalGetBaseMethod == nil {
        method = defaultGetBaseMethod
    }
    baseCtx := createBaseInfoContext(method)
    return createGlobalCacheContext(baseCtx)
}

func CreateTracingContextWithBaseStructure( operationName string, opts ... opentracing.StartSpanOption) context.Context {
    baseCtx := InitBasementContext()
    return CreateContextWithTracing(baseCtx, operationName, opts ...)
}