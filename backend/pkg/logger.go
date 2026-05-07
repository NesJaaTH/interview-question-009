package pkg

import (
	"log/slog"
	"os"
)

// InitLogger configures the global structured logger based on the environment.
// Production uses JSON output; development uses human-readable text.
func InitLogger(env string) {
	level := slog.LevelInfo
	if env != "production" {
		level = slog.LevelDebug
	}
	opts := &slog.HandlerOptions{Level: level}
	var handler slog.Handler
	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}
	slog.SetDefault(slog.New(handler))
}
