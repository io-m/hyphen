package helpers

import (
	"log/slog"
	"os"
	"time"
)

func InitStructuredLogger() {
	loggerHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "date"
				a.Value = slog.Int64Value(time.Now().UnixMilli())
			}
			return a
		},
	})
	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)
}
