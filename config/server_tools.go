package config

import (
	"github.com/hugoluchessi/gobservable/logging"
)

type ServerTools struct {
	Logger *logging.ContextLogger
}

func NewServerTools(l *logging.ContextLogger) *ServerTools {
	return &ServerTools{l}
}
