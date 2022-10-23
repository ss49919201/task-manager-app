package log

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	logger   zerolog.Logger
	syncOnce sync.Once
)

func setup() {
	syncOnce.Do(func() {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	})
}

func Fatal() *zerolog.Event {
	setup()
	return logger.Fatal()
}

func Error() *zerolog.Event {
	setup()
	return logger.Error()
}

func Warn() *zerolog.Event {
	setup()
	return logger.Warn()
}

func Debug() *zerolog.Event {
	setup()
	return logger.Debug()
}

func Info() *zerolog.Event {
	setup()
	return logger.Info()
}
