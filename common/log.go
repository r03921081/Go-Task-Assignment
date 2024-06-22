package common

import (
	"context"

	"github.com/andyliao/task-homework/constant"
	"github.com/sirupsen/logrus"

	"github.com/andyliao/task-homework/util"
)

type ILogger interface {
	Info(ctx context.Context, message string)
	Debug(ctx context.Context, message string)
	Warn(ctx context.Context, message string)
	Error(ctx context.Context, message string)
}

func InitLogger() {
	logrus.SetLevel(logrus.DebugLevel)
}

func GetLogger() *logrus.Logger {
	return logrus.StandardLogger()
}

type LoggerImpl struct{}

var Logger ILogger = &LoggerImpl{}

func (l *LoggerImpl) Info(ctx context.Context, message string) {
	GetLogger().WithFields(logrus.Fields{
		constant.HeaderTraceID: util.GetValueFromContext(ctx, constant.HeaderTraceID),
		constant.HeaderUserID:  util.GetValueFromContext(ctx, constant.HeaderUserID),
	}).Info(message)
}

func (l *LoggerImpl) Debug(ctx context.Context, message string) {
	GetLogger().WithFields(logrus.Fields{
		constant.HeaderTraceID: util.GetValueFromContext(ctx, constant.HeaderTraceID),
		constant.HeaderUserID:  util.GetValueFromContext(ctx, constant.HeaderUserID),
	}).Debug(message)
}

func (l *LoggerImpl) Warn(ctx context.Context, message string) {
	GetLogger().WithFields(logrus.Fields{
		constant.HeaderTraceID: util.GetValueFromContext(ctx, constant.HeaderTraceID),
		constant.HeaderUserID:  util.GetValueFromContext(ctx, constant.HeaderUserID),
	}).Warn(message)
}

func (l *LoggerImpl) Error(ctx context.Context, message string) {
	GetLogger().WithFields(logrus.Fields{
		constant.HeaderTraceID: util.GetValueFromContext(ctx, constant.HeaderTraceID),
		constant.HeaderUserID:  util.GetValueFromContext(ctx, constant.HeaderUserID),
	}).Error(message)
}
