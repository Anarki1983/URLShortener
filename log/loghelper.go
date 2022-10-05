package log

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type Fields logrus.Fields

func Init() {
	formatter := &prefixed.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000000",
		FullTimestamp:   true,
		ForceFormatting: true,
		ForceColors:     true,
	}

	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "34",
		WarnLevelStyle:  "11",
		ErrorLevelStyle: "9",
		FatalLevelStyle: "9",
		PanicLevelStyle: "9",
		DebugLevelStyle: "14",
		PrefixStyle:     "225",
		TimestampStyle:  "8",
	})

	logrus.SetFormatter(formatter)
	logrus.SetOutput(os.Stdout)
	SetLevel("Info")
}

func SetLevel(level string) {
	loglevel, err := logrus.ParseLevel(level)
	if err != nil {
		loglevel = logrus.InfoLevel
	}

	logrus.SetLevel(loglevel)
	Info(context.Background(), "set log level to [%s]", loglevel)
}

func decorateEntryWithCtx(ctx context.Context) *logrus.Entry {
	entry := logrus.NewEntry(logrus.StandardLogger())
	entry = decorateRuntimeContext(ctx, entry)

	return entry
}

func decorateRuntimeContext(ctx context.Context, logger *logrus.Entry) *logrus.Entry {
	if _, file, line, ok := runtime.Caller(3); ok {
		return logger.WithField("prefix", fmt.Sprintf("%s:%d", filepath.Base(file), line))
	}

	return logger
}

func Trace(ctx context.Context, format string, args ...interface{}) {
	decorateEntryWithCtx(ctx).Tracef(format, args...)
}

func Debug(ctx context.Context, format string, args ...interface{}) {
	decorateEntryWithCtx(ctx).Debugf(format, args...)
}

func Info(ctx context.Context, format string, args ...interface{}) {
	decorateEntryWithCtx(ctx).Infof(format, args...)
}

func Warn(ctx context.Context, format string, args ...interface{}) {
	decorateEntryWithCtx(ctx).Warnf(format, args...)
}

func Error(ctx context.Context, format string, args ...interface{}) {
	decorateEntryWithCtx(ctx).Errorf(format, args...)
}

func Fatal(ctx context.Context, format string, args ...interface{}) {
	decorateEntryWithCtx(ctx).Fatalf(format, args...)
}

func Panic(ctx context.Context, format string, args ...interface{}) {
	decorateEntryWithCtx(ctx).Panicf(format, args...)
}
