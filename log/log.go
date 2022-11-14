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

func lazyInit() {
	syncOnce.Do(func() {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	})
}

func setup(callback func() *zerolog.Event) *zerolog.Event {
	lazyInit()
	return callback()
}

func Fatal() *zerolog.Event {
	return setup(logger.Fatal)
}

func Error() *zerolog.Event {
	return setup(logger.Error)
}

func Warn() *zerolog.Event {
	return setup(logger.Warn)
}

func Debug() *zerolog.Event {
	return setup(logger.Debug)
}

func Info() *zerolog.Event {
	return setup(logger.Info)
}
