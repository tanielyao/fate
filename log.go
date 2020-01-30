package fate

import (
	"github.com/goextension/log"
	zap2 "github.com/goextension/log/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	cfg := zap.NewProductionConfig()

	cfg.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "",
		TimeKey:        "",
		NameKey:        "",
		CallerKey:      "",
		StacktraceKey:  "",
		LineEnding:     "",
		EncodeLevel:    nil,
		EncodeTime:     nil,
		EncodeDuration: nil,
		EncodeCaller:   nil,
		EncodeName:     nil,
	}
	cfg.OutputPaths = []string{
		"name.txt",
	}

	cfg.ErrorOutputPaths = []string{
		zap2.DefaultZapFilePath,
	}

	cfg.DisableCaller = true
	cfg.DisableStacktrace = true

	logger, e := cfg.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if e != nil {
		panic(e)
	}
	log.Register(logger.Sugar())
}
