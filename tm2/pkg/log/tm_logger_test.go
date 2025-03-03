package log_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/gnolang/gno/tm2/pkg/log"
)

func TestLoggerLogsItsErrors(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	logger := log.NewTMLogger(&buf)
	logger.Info("foo", "baz baz", "bar")
	msg := strings.TrimSpace(buf.String())
	if !strings.Contains(msg, "foo") {
		t.Errorf("Expected logger msg to contain foo, got %s", msg)
	}
}

func BenchmarkTMLoggerSimple(b *testing.B) {
	benchmarkRunner(b, log.NewTMLogger(io.Discard), baseInfoMessage)
}

func BenchmarkTMLoggerContextual(b *testing.B) {
	benchmarkRunner(b, log.NewTMLogger(io.Discard), withInfoMessage)
}

func benchmarkRunner(b *testing.B, logger log.Logger, f func(log.Logger)) {
	b.Helper()

	lc := logger.With("common_key", "common_value")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(lc)
	}
}

var (
	baseInfoMessage = func(logger log.Logger) { logger.Info("foo_message", "foo_key", "foo_value") }
	withInfoMessage = func(logger log.Logger) { logger.With("a", "b").Info("c", "d", "f") }
)
