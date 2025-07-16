package logging

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"
	"time"
)

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"
	msg := r.Message

	now := time.Now()
	timeStr := darkGray.Color(now.Format("[15:04:05]"))

	fields := make(map[string]any, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		if a.Value.Kind() == slog.KindAny {
			fields[a.Key] = a.Value.String()
		} else {
			fields[a.Key] = a.Value.Any()
		}
		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}
	switch r.Level {
	case slog.LevelDebug:
		level = gray.Color(level)
		msg = gray.Color(msg)
	case slog.LevelInfo:
		level = cyan.Color(level)
		msg = cyan.Color(msg)
	case slog.LevelWarn:
		level = yellow.Color(level)
		msg = yellow.Color(msg)
	case slog.LevelError:
		level = red.Color(level)
		msg = red.Color(msg)
	}

	data := string(b)
	if data == "{}" {
		h.l.Println(timeStr, level, msg)
	} else {
		h.l.Println(timeStr, level, msg, gray.Color(data))
	}
	return nil
}

func NewPrettyHandler(
	out io.Writer,
	opts *slog.HandlerOptions,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts),
		l:       log.New(out, "", 0),
	}

	return h
}

func SetPrettyDebugLogger() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	h := NewPrettyHandler(os.Stderr, opts)
	logger := slog.New(h)
	slog.SetDefault(logger)
}
