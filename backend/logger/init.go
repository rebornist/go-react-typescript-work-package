package logger

import "github.com/sirupsen/logrus"

// 로그 구조체 정의
type Logger struct {
	*logrus.Logger
	*logrus.Entry
}

// 생성자 선언
func NewLogger() LoggerService {
	logger := logrus.New()
	logEntry := logrus.NewEntry(logger)
	return &Logger{Logger: logger, Entry: logEntry}
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.InfoLevel)
}
