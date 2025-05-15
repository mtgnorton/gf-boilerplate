// Package st Prometheus初始化
package st

import (
	"context"

	"go.opentelemetry.io/otel/exporters/prometheus"

	"github.com/gogf/gf/contrib/metric/otelmetric/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gmetric"
)

func MustInitPrometheusByConfig(ctx context.Context, server *ghttp.Server) gmetric.Provider {
	if !g.Config().MustGet(ctx, "prometheus.enabled").Bool() {
		return nil
	}
	// 创建Prometheus导出器，用于指标数据的导出
	exporter, err := prometheus.New(
		prometheus.WithoutCounterSuffixes(), // Remove counter suffixes for cleaner metric names
		prometheus.WithoutUnits(),           // Remove unit suffixes for cleaner metric names
	)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	// 初始化并配置OpenTelemetry提供者
	provider := otelmetric.MustProvider(
		otelmetric.WithReader(exporter), // 配置指标读取器
		otelmetric.WithBuiltInMetrics(), // 启用内置指标收集
	)
	provider.SetAsGlobal() // 设置为全局指标提供者

	path := g.Config().MustGet(ctx, "prometheus.path").String()
	server.BindHandler(path, otelmetric.PrometheusHandler)
	return provider
}
