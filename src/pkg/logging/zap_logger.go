package logging

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logLevelMap = map[string]zapcore.Level{
	"debug":   zapcore.DebugLevel,
	"info":    zapcore.InfoLevel,
	"warning": zapcore.WarnLevel,
	"error":   zapcore.ErrorLevel,
	"fatal":   zapcore.FatalLevel,
}

type ZapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *config.Config) *ZapLogger {
	logger := &ZapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (z *ZapLogger) getLogLevel() zapcore.Level {
	level, ok := logLevelMap[z.cfg.Logger.Level]
	if !ok {
		return zapcore.InfoLevel
	}
	return level
}

func (z *ZapLogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   z.cfg.Logger.FilePath,
		MaxSize:    1,
		MaxAge:     5,
		Compress:   true,
		LocalTime:  true,
		MaxBackups: 10,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		z.getLogLevel(),
	)
	z.logger = zap.New(core).Sugar()

	z.logger = zap.New(core, zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
}

func (z *ZapLogger) Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	pairs := getPairs(extra, category, subCategory)
	z.logger.Debugw(message, pairs...)
}

func (z *ZapLogger) Debugf(template string, args ...interface{}) {
	z.logger.Debugf(template, args...)
}

func (z *ZapLogger) Info(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	pairs := getPairs(extra, category, subCategory)
	z.logger.Infow(message, pairs...)
}
func (z *ZapLogger) Infof(template string, args ...interface{}) {
	z.logger.Infof(template, args...)
}

func (z *ZapLogger) Warning(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	pairs := getPairs(extra, category, subCategory)
	z.logger.Warnw(message, pairs...)
}
func (z *ZapLogger) Warningf(template string, args ...interface{}) {
	z.logger.Warnf(template, args...)
}

func (z *ZapLogger) Error(err error, category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	pairs := getPairs(extra, category, subCategory)
	z.logger.Errorw(message, append(pairs, "error", err)...)
}
func (z *ZapLogger) Errorf(err error, template string, args ...interface{}) {
	z.logger.Errorf(template, args...)
}

func (z *ZapLogger) Fatal(err error, category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	pairs := getPairs(extra, category, subCategory)
	z.logger.Fatalw(message, append(pairs, "error", err)...)
}
func (z *ZapLogger) Fatalf(err error, template string, args ...interface{}) {
	z.logger.Fatalf(template, args...)
}

func getPairs(extra map[ExtraKey]interface{}, category Category, subCategory SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{})
	}
	extra["Category"] = category
	extra["SubCategory"] = subCategory
	pairs := ConvertMapToInterface(extra)
	return pairs
}
