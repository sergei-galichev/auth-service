package logging

import (
	"auth-service/internal/config/env"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"runtime"
)

// writerHook implements logrus.Hook.
// It is a hook that writes logs of specified LogLevels to specified Writer
type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Fire will be called when some logging function is called with current hook.
// It will format logbook entry to string and write it to appropriate writer
func (w *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range w.Writer {
		_, err = w.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	return err
}

// Levels define on which logbook levels this hook will be triggered
func (w *writerHook) Levels() []logrus.Level {
	return w.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

// Init initializes the logger
func Init() {
	cfg, err := env.NewLoggingConfig()
	if err != nil {
		log.Fatal("Failed to create logging config: ", err)
	}

	var level logrus.Level

	switch cfg.LoggingLevel() {
	case "info":
		level = logrus.InfoLevel
	case "warn":
		level = logrus.WarnLevel
	case "error":
		level = logrus.ErrorLevel
	case "debug":
		level = logrus.DebugLevel
	default:
		level = logrus.TraceLevel
	}
	_ = level

	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s:%d", filename, frame.Line), fmt.Sprintf("%s()", frame.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	makeDirErr := os.MkdirAll("logs", 0755)

	if makeDirErr != nil || os.IsExist(makeDirErr) {
		log.Fatal("Failed to create logs directory")
	} else {
		allLogsFile, openErr := os.OpenFile("logs/all.logbook", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if openErr != nil {
			panic(fmt.Sprintf("[Message]: %s", openErr))
		}

		l.SetOutput(io.Discard)
		l.AddHook(
			&writerHook{
				Writer:    []io.Writer{allLogsFile, os.Stdout},
				LogLevels: logrus.AllLevels,
			},
		)
	}
	e = logrus.NewEntry(l)
	e.Logger.SetLevel(level)
}

func GetLogger() *Logger {
	return &Logger{
		Entry: e,
	}
}

func (l *Logger) GetLoggerWithField(key string, value interface{}) *Logger {
	return &Logger{
		l.WithField(key, value),
	}
}
