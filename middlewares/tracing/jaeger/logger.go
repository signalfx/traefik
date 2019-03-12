package jaeger

// From https://github.com/containous/traefik/blob/156f6b8d3c62ab6df1b46f656e672498f5057353/tracing/jaeger/logger.go
import (
	"github.com/containous/traefik/log"
	"github.com/sirupsen/logrus"
)

// jaegerLogger is an implementation of the Logger interface that delegates to traefik log
type jaegerLogger struct {
	logger logrus.FieldLogger
}

func newJaegerLogger() *jaegerLogger {
	return &jaegerLogger{
		logger: log.WithoutContext().WithField(log.TracingProviderName, "jaeger"),
	}
}

func (l *jaegerLogger) Error(msg string) {
	l.logger.Errorf("Tracing jaeger error: %s", msg)
}

// Infof logs a message at debug priority
func (l *jaegerLogger) Infof(msg string, args ...interface{}) {
	l.logger.Debugf(msg, args...)
}
