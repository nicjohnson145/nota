package report

import (
	"github.com/rs/zerolog"
)

type LogReporterConfig struct {
	Logger zerolog.Logger
}

func NewLogReporter(conf LogReporterConfig) *LogReporter {
	return &LogReporter{
		log: conf.Logger,
	}
}

var _ Reporter = (*LogReporter)(nil)

type LogReporter struct {
	log zerolog.Logger
}

func (l *LogReporter) Report(msg string) {
	l.log.Error().Msg(msg)
}
