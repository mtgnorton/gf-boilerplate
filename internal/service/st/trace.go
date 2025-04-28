// Package st 全局单例-灵活配置
package st

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	traceOption "go.opentelemetry.io/otel/trace"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/net/gtrace"
)

const (
	tracerHostnameTagKey = "hostname"
)

var defaultTracer = &Tracer{}

// Tracer 链路追踪器
type Tracer struct {
}

// GetTracer 获取链路追踪器
func GetTracer() *Tracer {
	return defaultTracer
}

func (t *Tracer) NewSpan(ctx context.Context, spanName string, opts ...traceOption.SpanStartOption) (context.Context, *gtrace.Span) {
	if !g.Cfg().MustGet(ctx, "trace.enabled").Bool() {
		return ctx, &gtrace.Span{
			Span: noopSpanInstance,
		}
	}
	ctx, span := gtrace.NewTracer().Start(ctx, spanName, opts...)
	return ctx, &gtrace.Span{
		Span: span,
	}
}

// MustInitByConfig 初始化链路追踪器,如果出错直接panic
func (t *Tracer) MustInitByConfig(ctx context.Context) func(ctx context.Context) {
	if !g.Cfg().MustGet(ctx, "trace.enabled").Bool() {
		return func(ctx context.Context) {}
	}
	serviceName := g.Cfg().MustGet(ctx, "trace.serviceName").String()
	endpoint := g.Cfg().MustGet(ctx, "trace.endpoint").String()
	path := g.Cfg().MustGet(ctx, "trace.path").String()
	cleanup, err := t.Init(serviceName, endpoint, path)
	if err != nil {
		g.Log().Panicf(ctx, "初始化链路追踪器失败 err:%+v", err)
	}
	return cleanup
}

// Init 初始化链路追踪器
//
// 参数说明：
//   serviceName: 服务名称，在追踪后端（如 Jaeger、Zipkin、Tempo 等）中用于区分不同服务，便于检索和展示。
//   endpoint: OTLP HTTP 导出器的服务端地址（如 "localhost:4318"），用于指定 trace 数据上报的目标后端。
//   path: OTLP HTTP 导出器的 URL 路径（如 "/v1/traces"），用于指定 trace 数据上报的具体接口路径。
//
// 返回值说明:
//   - func(ctx context.Context): 清理函数,用于关闭追踪器
//   - error: 初始化过程中的错误信息
//
// 注意事项:
//   - 需要确保 endpoint 地址可访问
//   - 建议在程序退出时调用返回的清理函数
//
// 使用示例:
//
//	cleanup, err := TracerInit("myapp", "http://localhost:14268/api/traces", "test")
//	if err != nil {
//	    panic(err)
//	}
//	defer cleanup(context.Background())

func (t *Tracer) Init(serviceName, endpoint, path string) (func(ctx context.Context), error) {
	// 尝试获取主机 IP 用于追踪信息
	var (
		intranetIPArray, err = gipv4.GetIntranetIpArray()
		hostIP               = "NoHostIpFound"
	)

	if err != nil {
		return nil, err
	}

	// 如果获取内网 IP 失败,则尝试获取所有 IP
	if len(intranetIPArray) == 0 {
		if intranetIPArray, err = gipv4.GetIpArray(); err != nil {
			return nil, err
		}
	}
	if len(intranetIPArray) > 0 {
		hostIP = intranetIPArray[0]
	}

	ctx := context.Background()
	// 创建 OTLP HTTP 导出器
	traceExp, err := otlptrace.New(ctx, otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithURLPath(path),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithCompression(1),
	))
	if err != nil {
		return nil, err
	}

	// 创建资源属性
	res, err := resource.New(ctx,
		resource.WithFromEnv(),      // 从环境变量中获取资源属性
		resource.WithProcess(),      // 添加进程信息
		resource.WithTelemetrySDK(), // 添加遥测SDK信息
		resource.WithHost(),         // 添加主机信息
		resource.WithAttributes(
			// 在追踪后端显示的服务名称
			semconv.ServiceNameKey.String(serviceName),
			semconv.HostNameKey.String(hostIP),
			attribute.String(tracerHostnameTagKey, hostIP),
		),
	)

	// 创建并配置 TracerProvider
	tracerProvider := trace.NewTracerProvider(
		// 使用 AlwaysSample 采样器,对所有 trace 进行采样
		trace.WithSampler(trace.AlwaysSample()),
		// 设置资源属性
		trace.WithResource(res),
		// 设置批处理 span 处理器
		trace.WithSpanProcessor(trace.NewBatchSpanProcessor(traceExp)),
	)

	// 设置全局传播器为 TraceContext (默认未设置)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)
	otel.SetTracerProvider(tracerProvider)

	g.Log().Infof(ctx, "初始化链路追踪器成功 serviceName:%s endpoint:%s path:%s", serviceName, endpoint, path)
	// 返回关闭函数
	return func(ctx context.Context) {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		if err = tracerProvider.Shutdown(ctx); err != nil {
			g.Log().Errorf(ctx, "关闭 tracerProvider 失败 err:%+v", err)
		} else {
			g.Log().Debug(ctx, "关闭 tracerProvider 成功")
		}
	}, nil
}
