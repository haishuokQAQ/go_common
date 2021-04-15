package tracing

import (
    `context`
    `github.com/opentracing/opentracing-go`
    `testing`
    `time`
)

func TestSendSpan(t *testing.T) {
    err := InitJaegerTracer("test-service", "http://127.0.0.1:14268/api/traces")
    if err != nil {
        panic(err)
    }
    span := opentracing.StartSpan("test-option")
    span.Finish()
    spanContext := opentracing.ContextWithSpan(context.Background(), span)
    span1, newCtx := opentracing.StartSpanFromContext(spanContext, "test-option2")
    span1.Finish()
    newCtx.Done()
    time.Sleep(1 * time.Minute)
}
