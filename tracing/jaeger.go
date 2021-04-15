package tracing

import (
    `github.com/uber/jaeger-client-go`
    jaegercfg "github.com/uber/jaeger-client-go/config"
    jaegerlog "github.com/uber/jaeger-client-go/log"
)

func InitJaegerTracer(serviceName string, remoteAddr string) error {
    conf := &jaegercfg.Configuration{
        ServiceName:         serviceName,
        Tags:                nil,
        Sampler:             &jaegercfg.SamplerConfig{
            Type:                     jaeger.SamplerTypeConst,
            Param:                    1,
        },
        Reporter:            &jaegercfg.ReporterConfig{
            BufferFlushInterval:        1,
            LogSpans:                   true,
            CollectorEndpoint: remoteAddr,
        },
    }
    _, err := conf.InitGlobalTracer(
        serviceName,
        jaegercfg.Logger(jaegerlog.StdLogger),
    )
    if err != nil {
        return err
    }
    return nil
}
