package ovirtclientlog_test

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/ovirt/go-ovirt-client-log/v2"
)

func TestGoLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	backingLogger := log.New(buf, "", 0)
	logger := ovirtclientlog.NewGoLogger(backingLogger)
	logger.Infof("Hello world!")
	if !strings.Contains(buf.String(), "Hello world!\n") {
		t.Fatalf("failed to find written log message in output")
	}
}
