package context_strengthen

import (
    `context`
    `github.com/opentracing/opentracing-go`
)

// 使用opentracing
func CreateContextWithTracing(parent context.Context, operationName string, opts ... opentracing.StartSpanOption) context.Context {
    _, ctx := opentracing.StartSpanFromContext(parent, operationName, opts...)
    return ctx
}

func GetSpanFromContext(ctx context.Context) opentracing.Span {
    return opentracing.SpanFromContext(ctx)
}

