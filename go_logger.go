package ovirtclientlog

import (
	"fmt"
	"log"
)

// NewGoLogger creates a logger that writes to the Go log facility. The optional logger parameter can be
// used to create a scoped logger, otherwise the global logger is used.
func NewGoLogger(logger *log.Logger) Logger {
	return &goLogger{
		logger: logger,
	}
}

type goLogger struct {
	logger *log.Logger
}

func (g *goLogger) write(format string, args ...interface{}) {
	if g.logger == nil {
		log.Printf(fmt.Sprintf("%s\n", format), args...)
	} else {
		g.logger.Printf(fmt.Sprintf("%s\n", format), args...)
	}
}

func (g *goLogger) Debugf(format string, args ...interface{}) {
	g.write(format, args...)
}

func (g *goLogger) Infof(format string, args ...interface{}) {
	g.write(format, args...)
}

func (g *goLogger) Warningf(format string, args ...interface{}) {
	g.write(format, args...)
}

func (g *goLogger) Errorf(format string, args ...interface{}) {
	g.write(format, args...)
}
