package main

import (
	"log/slog"
	"os"

	"github.com/yanmoyy/go-go-go/internal/logging"
)

func main() {
	opts := logging.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := logging.NewPrettyHandler(os.Stdout, opts)
	logger := slog.New(handler)
	logger.Debug(
		"executing database query",
		slog.String("query", "SELECT * FROM users"),
	)
	logger.Info("image upload successful", slog.String("image_id", "39ud88"))
	logger.Warn(
		"storage is 90% full",
		slog.String("available_space", "900.1 MB"),
	)
	logger.Error(
		"An error occurred while processing the request",
		slog.String("url", "https://example.com"),
	)
}
