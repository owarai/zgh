package log

import (
	"context"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DefaultFileName = "go_blog"
	// logKey is the key for log.Logger values in Contexts. It is
	// unexported; clients use log.NewContext and log.FromContext
	// instead of using this key directly.
	logKey ctxKey = 0
)

type (
	Logger = *zap.SugaredLogger
	// ctxKey is an unexported type for keys defined in this package.
	// This prevents collisions with keys defined in other packages.
	ctxKey int
)

func L() Logger {
	return zap.S()
}

// New create a logger which write to dest: logDir + appName + ".log"/"error.log"
//
// This function also initializes the package-level variable logger, which is used for GetEventLoggerByReqId.
func New(logDir string, appName string) Logger {
	destWithoutExt := filepath.Join(logDir, appName)
	logger := newZapLog(destWithoutExt)
	_ = zap.ReplaceGlobals(logger.Desugar())
	return logger
}

func newZapLog(destName string) Logger {
	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.
	logDebugs := zapcore.AddSync(&lumberjack.Logger{
		Filename:   destName + ".log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     14, // days
	})
	logErrors := zapcore.AddSync(&lumberjack.Logger{
		Filename:   destName + ".error.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     14, // days
	})

	// Optimize the log output for machine consumption and the console output
	// for human operators.
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = defaultTimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, logErrors, highPriority),
		zapcore.NewCore(encoder, logDebugs, zapcore.DebugLevel),
	)

	return zap.New(core).Sugar()
}

// NewContext returns a new Context that carries value l.
func NewContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, logKey, l)
}

// FromContext returns the Logger value stored in ctx, if any.
func FromContext(ctx context.Context) (Logger, bool) {
	l, ok := ctx.Value(logKey).(Logger)
	return l, ok
}
