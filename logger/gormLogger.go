/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 09:51:05
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 10:32:55
 * @FilePath: \ginManager\logger\gormLogger.go
 * @Description: GORM 日志适配 zap
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */

package logger

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// zapLogger 实现 gorm.Logger 接口
type zapLogger struct {
	zap           *zap.Logger
	level         zapcore.Level
	slowThreshold time.Duration
}

// NewGormLogger 构造函数，供外部调用
func NewGormLogger(z *zap.Logger) logger.Interface {
	return &zapLogger{
		zap:           z,
		level:         zapcore.InfoLevel, // 默认 Info
		slowThreshold: 200 * time.Millisecond,
	}
}

// LogMode 实现接口：调整日志级别
func (l *zapLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	switch level {
	case logger.Silent:
		newLogger.level = zapcore.FatalLevel + 1 // 比 Fatal 还高，基本不输出
	case logger.Error:
		newLogger.level = zapcore.ErrorLevel
	case logger.Warn:
		newLogger.level = zapcore.WarnLevel
	case logger.Info:
		newLogger.level = zapcore.InfoLevel
	}
	return &newLogger
}

// Info 打印 SQL 以外的信息
func (l *zapLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level <= zapcore.InfoLevel {
		l.zap.Sugar().Infof(msg, data...)
	}
}

// Warn 打印警告
func (l *zapLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level <= zapcore.WarnLevel {
		l.zap.Sugar().Warnf(msg, data...)
	}
}

// Error 打印错误
func (l *zapLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level <= zapcore.ErrorLevel {
		l.zap.Sugar().Errorf(msg, data...)
	}
}

// Trace 打印 SQL 语句
func (l *zapLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := []zapcore.Field{
		zap.String("sql", sql),
		zap.Duration("elapsed", elapsed),
		zap.Int64("rows", rows),
	}

	switch {
	case err != nil && !errors.Is(err, gorm.ErrRecordNotFound):
		l.zap.Error("gorm error", append(fields, zap.Error(err))...)
	case elapsed > l.slowThreshold:
		l.zap.Warn("gorm slow sql", fields...)
	default:
		if l.level <= zapcore.InfoLevel {
			l.zap.Info("gorm trace", fields...)
		}
	}
}
