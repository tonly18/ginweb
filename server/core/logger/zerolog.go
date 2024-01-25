package logger

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"runtime"
	"server/config"
	"server/core/logger/hook"
	"time"
)

// logger
var logger zerolog.Logger

// init
func init() {
	fs, err := os.OpenFile(config.Config.Log.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprint(`open log file is error: %w`, err))
	}
	runtime.SetFinalizer(fs, func(f *os.File) {
		f.Close()
	})

	zerolog.TimeFieldFormat = time.DateTime
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger = zerolog.New(fs).With().Timestamp().Logger().Hook(&hook.ZeroLogHook{})
}

// Debug
func Debug(ctx context.Context, args ...any) {
	if ctx == nil {
		ctx = context.Background()
	}
	logger.Debug().Fields(map[string]any{
		"trace_id": ctx.Value("trace_id"),
		"user_id":  ctx.Value("user_id"),
		"ip":       ctx.Value("client_ip"),
	}).Msg(fmt.Sprint(args...))
}
func Debugf(ctx context.Context, format string, args ...any) {
	if ctx == nil {
		ctx = context.Background()
	}
	logger.Debug().Fields(map[string]any{
		"trace_id": ctx.Value("trace_id"),
		"user_id":  ctx.Value("user_id"),
		"ip":       ctx.Value("client_ip"),
	}).Msgf(format, args...)
}

// Info
func Info(ctx context.Context, args ...any) {
	if ctx == nil {
		ctx = context.Background()
	}
	logger.Info().Fields(map[string]any{
		"trace_id": ctx.Value("trace_id"),
		"user_id":  ctx.Value("user_id"),
		"ip":       ctx.Value("client_ip"),
	}).Msg(fmt.Sprint(args...))
}
func Infof(ctx context.Context, format string, args ...any) {
	if ctx == nil {
		ctx = context.Background()
	}
	logger.Info().Fields(map[string]any{
		"trace_id": ctx.Value("trace_id"),
		"user_id":  ctx.Value("user_id"),
		"ip":       ctx.Value("client_ip"),
	}).Msgf(format, args...)
}

// Warning
func Warning(ctx context.Context, args ...any) {
	if ctx == nil {
		ctx = context.Background()
	}
	logger.Warn().Fields(map[string]any{
		"trace_id": ctx.Value("trace_id"),
		"user_id":  ctx.Value("user_id"),
		"ip":       ctx.Value("client_ip"),
	}).Msg(fmt.Sprint(args...))

}
func Warningf(ctx context.Context, format string, args ...any) {
	if ctx == nil {
		ctx = context.Background()
	}
	logger.Warn().Fields(map[string]any{
		"trace_id": ctx.Value("trace_id"),
		"user_id":  ctx.Value("user_id"),
		"ip":       ctx.Value("client_ip"),
	}).Msgf(format, args...)
}

// Error
func Error(ctx context.Context, args ...any) {
	if ctx == nil {
		ctx = context.Background()
	}
	logger.Error().Fields(map[string]any{
		"trace_id": ctx.Value("trace_id"),
		"user_id":  ctx.Value("user_id"),
		"ip":       ctx.Value("client_ip"),
	}).Msg(fmt.Sprint(args...))
}
func Errorf(ctx context.Context, format string, args ...any) {
	if ctx == nil {
		ctx = context.Background()
	}
	logger.Error().Fields(map[string]any{
		"trace_id": ctx.Value("trace_id"),
		"user_id":  ctx.Value("user_id"),
		"ip":       ctx.Value("client_ip"),
	}).Msgf(format, args...)
}
