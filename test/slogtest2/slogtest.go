package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	encoding := "console"
	outputPaths := []string{"stdout"}
	errorOutputPaths := []string{"stderr"}
	configs := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.DebugLevel), // 日志级别
		DisableStacktrace: true,
		Development:       false,                             // 开发模式，堆栈跟踪
		Encoding:          encoding,                          // 输出格式 console 或 json
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(), // 编码器配置
		InitialFields: map[string]interface{}{
			"service": "slogtest2",
		}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      outputPaths, // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: errorOutputPaths,
	}
	options := make([]zap.Option, 2)
	options[0] = zap.AddCallerSkip(1)
	options[1] = zap.AddStacktrace(zapcore.InfoLevel)

	logger, err := configs.Build(options...)
	if err != nil {
		panic(err)
	}
	// zaplog := logger.Sugar()
	// logger.AddCallerSkip()

	s := slog.New(NewZapHandler(logger))
	slog.SetDefault(s)

	// test
	calledFunc(logger)

}

func calledFunc(zaplog *zap.Logger) {
	ctx := context.Background()
	errDemo := errors.New("Oh my God!!!")
	slog.ErrorContext(ctx, "msg", errDemo)
	fmt.Println("*******************************************************")
	zaplog.Error("msg", zap.Error(errDemo))
}

var levelMap = map[slog.Level]zapcore.Level{
	slog.LevelDebug: zap.DebugLevel,
	slog.LevelInfo:  zap.InfoLevel,
	slog.LevelWarn:  zap.WarnLevel,
	slog.LevelError: zap.ErrorLevel,
}

var levelMap2 = map[zapcore.Level]slog.Level{
	zap.DebugLevel:  slog.LevelDebug,
	zap.InfoLevel:   slog.LevelInfo,
	zap.WarnLevel:   slog.LevelWarn,
	zap.ErrorLevel:  slog.LevelError,
	zap.DPanicLevel: slog.LevelError,
	zap.PanicLevel:  slog.LevelError,
	zap.FatalLevel:  slog.LevelError,
}

func NewZapHandler(zapLogger *zap.Logger) *ZapHandler {
	level := levelMap2[zapLogger.Level()]
	return &ZapHandler{
		attrs:  make([]slog.Attr, 0),
		groups: make([]string, 0),
		level:  level,
		log:    zapLogger,
	}
}

type ZapHandler struct {
	// option Option
	attrs  []slog.Attr
	groups []string
	level  slog.Level
	log    *zap.Logger
}

func (h *ZapHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *ZapHandler) Handle(ctx context.Context, r slog.Record) error {
	var fields []zap.Field
	r.Attrs(func(attr slog.Attr) bool {
		key := attr.Key
		if key == "!BADKEY" {
			if _, ok := attr.Value.Any().(error); ok {
				key = "error"
			}
		}
		fields = append(fields, zap.Any(key, attr.Value))
		return true
	})

	level := levelMap[r.Level]
	h.log.Log(level, r.Message, fields...)

	return nil
}

func (h *ZapHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandler := *h
	for _, attr := range attrs {
		if attr.Key == "!BADKEY" {
			if _, ok := attr.Value.Any().(error); ok {
				attr.Key = "error"
			}
		}
		newHandler.attrs = append(newHandler.attrs, attr)
	}
	return &newHandler
}

func (h *ZapHandler) WithGroup(name string) slog.Handler {
	newHandler := *h
	newHandler.groups = append(newHandler.groups, name)
	return &newHandler
}
