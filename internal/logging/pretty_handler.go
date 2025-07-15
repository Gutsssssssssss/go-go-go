package logging

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"time"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = magenta.Color(level)
	case slog.LevelInfo:
		level = blue.Color(level)
	case slog.LevelWarn:
		level = yellow.Color(level)
	case slog.LevelError:
		level = red.Color(level)
	}

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

	now := time.Now()
	timeStr := gray.Color(now.Format("[15:04:05.000]"))

	msg := cyan.Color(r.Message)

	h.l.Println(timeStr, level, msg, white.Color(string(b)))

	return nil
}

func NewPrettyHandler(
	out io.Writer,
	opts PrettyHandlerOptions,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}
