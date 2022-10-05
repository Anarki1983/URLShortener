package log

import (
	"context"
	"testing"
)

func TestLog(t *testing.T) {
	Init()
	SetLevel("trace")

	ctx := context.Background()
	Trace(ctx, "trace log")
	Debug(ctx, "debug log")
	Info(ctx, "info log")
	Warn(ctx, "warn log")
	Error(ctx, "error log")
}
