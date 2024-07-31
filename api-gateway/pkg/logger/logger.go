package logger

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(logger *slog.Logger, appName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		// Process request
		ctx.Next()

		param := gin.LogFormatterParams{
			Request: ctx.Request,
			Keys:    ctx.Keys,
		}

		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = ctx.ClientIP()
		param.Method = ctx.Request.Method
		param.StatusCode = ctx.Writer.Status()
		param.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = ctx.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path
		event := slog.LevelInfo

		if param.StatusCode >= 400 && param.StatusCode < 500 {
			event = slog.LevelWarn
		}

		if param.StatusCode >= 500 {
			event = slog.LevelError
		}

		slog.LogAttrs(
			ctx.Request.Context(),
			event,
			param.ErrorMessage,
			slog.String("module", "gin"),
			slog.String("path", param.Path),
			slog.Int("status_code", param.StatusCode),
			slog.Float64("latency", float64(param.Latency)/float64(time.Millisecond)),
			slog.String("client_ip", param.ClientIP),
			slog.String("method", param.Method),
			slog.Int("body_size", param.BodySize),
		)
	}
}
