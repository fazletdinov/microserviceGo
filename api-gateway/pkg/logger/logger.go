package logger

import (
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type ginHands struct {
	Hostname   string
	StatusCode int
	Latency    int // time to process
	ClientIP   string
	UserAgent  string
	Method     string
	Path       string
	Referer    string
	DataLength int
	MsgStr     string
}

func ErrorLogger() gin.HandlerFunc {
	return ErrorLoggerT(gin.ErrorTypeAny)
}

func ErrorLoggerT(typ gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if !c.Writer.Written() {
			json := c.Errors.ByType(typ).JSON()
			if json != nil {
				c.JSON(-1, json)
			}
		}
	}
}

func Logger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		ctx.Next()

		if raw != "" {
			path = path + "?" + raw
		}
		msg := ctx.Errors.String()
		if msg == "" {
			msg = "Request"
		}
		stop := time.Since(start)

		cData := &ginHands{
			Hostname:   ctx.Request.Host,
			StatusCode: ctx.Writer.Status(),
			Latency:    int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0)),
			ClientIP:   ctx.ClientIP(),
			UserAgent:  ctx.Request.UserAgent(),
			Method:     ctx.Request.Method,
			Path:       path,
			Referer:    ctx.Request.Referer(),
			DataLength: ctx.Writer.Size(),
			MsgStr:     msg,
		}

		logSwitch(logger, cData)
	}
}
func logSwitch(log *zerolog.Logger, data *ginHands) {
	switch {
	case data.StatusCode >= 400 && data.StatusCode < 500:
		{
			log.Warn().
				Str("hostname", data.Hostname).
				Int("status", data.StatusCode).
				Int("latency", data.Latency).
				Str("client_ip", data.ClientIP).
				Str("method", data.Method).
				Str("path", data.Path).
				Str("referer", data.Referer).
				Int("data_length", data.DataLength).
				Msg(data.MsgStr)
		}
	case data.StatusCode >= 500:
		{
			log.Error().
				Str("hostname", data.Hostname).
				Int("status", data.StatusCode).
				Int("latency", data.Latency).
				Str("client_ip", data.ClientIP).
				Str("method", data.Method).
				Str("path", data.Path).
				Str("referer", data.Referer).
				Int("data_length", data.DataLength).
				Msg(data.MsgStr)
		}
	default:
		log.Info().
			Str("hostname", data.Hostname).
			Int("status", data.StatusCode).
			Int("latency", data.Latency).
			Str("client_ip", data.ClientIP).
			Str("method", data.Method).
			Str("path", data.Path).
			Str("referer", data.Referer).
			Int("data_length", data.DataLength).
			Msg(data.MsgStr)
	}
}
