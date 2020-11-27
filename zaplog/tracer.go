package zaplog

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

const (
	DefaultTraceIdName  = "trace_id"
	DefaultSpanIdName   = "span_id"
	DefaultParentIdName = "parent_id"
	// 默认分隔符
	DefaultSeparator = ":"
)

// LogTraceKey 定义日志中Trace的key
var LogTraceKey struct{} //格式 traceid:spanid:parentid  46b1506e7332f7c1:7f75737aa70629cc:3bb947500f42ad71:1

// ErrNoTracerInfo 定义错误
var ErrNoTracerInfo = errors.New("no trace info")

// GetTraceFields 获取链路跟踪添加列
func GetTraceFields(ctx context.Context) []zap.Field {
	fm := getTraceInfo(ctx)
	zf := make([]zap.Field, 0)
	for k, v := range fm {
		zf = append(zf, zap.String(k, v))
	}
	return zf
}

// getTraceInfo 定义日志中链路跟踪的信息
func getTraceInfo(ctx context.Context) map[string]string {
	s, err := decodeTracer(ctx)
	trace := make(map[string]string)
	if err != nil {
		return trace
	}
	trace[DefaultTraceIdName] = s[0]
	trace[DefaultSpanIdName] = s[1]
	trace[DefaultParentIdName] = s[2]
	return trace
}

// decodeTracer 解析trace中的信息
func decodeTracer(ctx context.Context) ([]string, error) {
	s := make([]string, 0, 4)
	if val, ok := ctx.Value(LogTraceKey).(string); ok {
		s = strings.Split(val, DefaultSeparator)
	} else {
		span := opentracing.SpanFromContext(ctx)
		s = strings.Split(fmt.Sprintf("%v", span), DefaultSeparator)
	}
	if len(s) >= 3 {
		return s, nil
	}
	return []string{}, ErrNoTracerInfo
}
