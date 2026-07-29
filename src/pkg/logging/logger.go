package logging

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
)

type Logger interface {
	Init()

	Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Info(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Warning(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Warningf(template string, args ...interface{})

	Error(err error, category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Errorf(err error, template string, args ...interface{})

	Fatal(err error, category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Fatalf(err error, template string, args ...interface{})
}

func NewLogger(cfg *config.Config) Logger {
	switch cfg.Logger.Logger {
	case "zap":
		return NewZapLogger(cfg)
	case "zero":
		return nil
	default:
		return NewZapLogger(cfg)
	}
}
