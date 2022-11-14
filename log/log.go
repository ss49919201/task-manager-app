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

func setupLogger(callback func() *zerolog.Event) *zerolog.Event {
	lazyInit()
	return callback()
}

func Fatal() *zerolog.Event {
	return setupLogger(logger.Fatal)
}

func Error() *zerolog.Event {
	return setupLogger(logger.Error)
}

func Warn() *zerolog.Event {
	return setupLogger(logger.Warn)
}

func Debug() *zerolog.Event {
	return setupLogger(logger.Debug)
}

func Info() *zerolog.Event {
	return setupLogger(logger.Info)
}
