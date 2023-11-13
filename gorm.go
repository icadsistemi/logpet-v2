package logpet

import (
	"context"
	"time"

	"gorm.io/gorm/logger"
)

func CreateGormLogger(logger *StandardLogger) *GormLogger {
	return &GormLogger{Logger: logger}
}

type GormLogger struct {
	Logger *StandardLogger
}

func (g GormLogger) LogMode(logger.LogLevel) logger.Interface {
	return g
}

func (g GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	g.Logger.SendDebugfLog("%s", nil, msg)
}
func (g GormLogger) Warn(context.Context, string, ...interface{}) {

}
func (g GormLogger) Error(context.Context, string, ...interface{}) {

}

func (g GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
