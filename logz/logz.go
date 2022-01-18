package logz

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Debug = "debug"
)

type Logging struct {
	Logger *zap.Logger
	level  zapcore.Level
	//TraceParent *traceparent.TraceParent
	ParentID string
}

func NewLogging(log *zap.Logger, level zapcore.Level) *Logging {
	return &Logging{
		Logger: log,
		level:  level,
		//ParentID:    parentID,
	}
}

func (l Logging) IsDebug() bool {
	return l.level.Enabled(zapcore.DebugLevel)
}

func Init(level zapcore.Level, format string) (*zap.Logger, error) {

	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.ISO8601TimeEncoder

	c := zap.NewProductionConfig()
	c.Level.SetLevel(level)
	c.EncoderConfig = ec
	c.Encoding = format
	c.DisableStacktrace = true

	return c.Build()
}

func ParseLevel(level string) zapcore.Level {
	switch level {
	default:
		return zapcore.InfoLevel
	case Debug:
		return zapcore.DebugLevel
	}
}
