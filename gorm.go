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

func (g GormLogger) LogMode(logger.LogLevel) {}

func (g GormLogger) Info(context.Context, string, ...interface{}) {

}
func (g GormLogger) Warn(context.Context, string, ...interface{}) {

}
func (g GormLogger) Error(context.Context, string, ...interface{}) {

}

func (g GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
