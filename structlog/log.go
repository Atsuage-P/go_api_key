package structlog

import (
	"context"
	"log/slog"
	"sync"
)

var _ slog.Handler = &LogHandler{}

type LogHandler struct {
	handler slog.Handler
}

func NewLogHandler(handler slog.Handler) slog.Handler {
	return LogHandler{
		handler: handler,
	}
}

func (h LogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h LogHandler) Handle(ctx context.Context, record slog.Record) error {
	if v, ok := ctx.Value(fields).(*sync.Map); ok {
		v.Range(func(key, val any) bool {
			if keyString, ok := key.(string); ok {
				record.AddAttrs(slog.Any(keyString, val))
			}
			return true
		})
	}
	return h.handler.Handle(ctx, record)
}

func (h LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return LogHandler{h.handler.WithAttrs(attrs)}
}

func (h LogHandler) WithGroup(name string) slog.Handler {
	return h.handler.WithGroup(name)
}
