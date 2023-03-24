package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// 로그 출력 기능
func (l *Logger) PrintLogger(logdata map[string]interface{}, status int, message string) error {

	// 화면에 출력하는 옵션
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// request 데이터를 log entry 값으로 변환
	log := l.Entry.WithFields(logrus.Fields{
		"backoff":     time.Second,
		"body":        logdata["body"],
		"created":     time.Now(),
		"IP":          logdata["connect_ip"],
		"message":     message,
		"request-id":  logdata["request_id"],
		"request-url": logdata["request_url"],
		"status":      status,
		"user-agent":  logdata["user_agent"],
	})

	// 에러 메시지 존재 여부에 따른 출력 방식 설정
	if message != "" {
		log.Error(message)
	} else {
		log.Info("")
	}

	return nil
}

// 로그 작성 기능
func (l *Logger) WriteLogger(logdata map[string]interface{}, status int, message string) error {

	// 해당 년도와 주를 출력
	year, week := time.Now().ISOWeek()

	// request 데이터를 log entry 값으로 변환
	log := l.Entry.WithFields(logrus.Fields{
		"backoff":     time.Second,
		"body":        logdata["body"],
		"created":     time.Now(),
		"IP":          logdata["connect_ip"],
		"message":     message,
		"request-id":  logdata["request_id"],
		"request-url": logdata["request_url"],
		"status":      status,
		"user-agent":  logdata["user_agent"],
	})

	// 로그 파일 저장 전 세팅
	var file *os.File
	var err error
	if message != "" {
		// 파일 생성
		file, err = os.OpenFile(fmt.Sprintf("log/error_%d-%d.log", year, week), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("%s : %s", "Failed to log to file", err.Error())
		}
	} else {
		// 파일 생성
		file, err = os.OpenFile(fmt.Sprintf("log/access_%d-%d.log", year, week), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("%s : %s", "Failed to log to file", err.Error())
		}
	}

	// 파일에 저장하는 옵션
	l.Logger.Out = file

	// json 형식으로 출력
	log.Logger.SetFormatter(&logrus.JSONFormatter{})

	// 파일 저장
	log.Info(message)

	return nil
}

// 로그 기능 에러 시 패닉 출력 후 리커버
func (l *Logger) RaiseError(err error) {
	// 리커버
	defer func() {
		if r := recover(); r != nil {
			// 에러 내역 출력
			fmt.Printf("Recover Logger : %s\n", err.Error())
		}
	}()

	// 해당 년도와 주를 출력
	year, week := time.Now().ISOWeek()

	// 파일 생성
	file, _ := os.OpenFile(fmt.Sprintf("panic_%d-%d.log", year, week), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// 파일에 저장하는 옵션
	l.Logger.Out = file

	// 에러 데이터를 log entry 값으로 변환
	log := l.Entry.WithFields(logrus.Fields{"message": err.Error()})

	// json 형식으로 출력
	log.Logger.SetFormatter(&logrus.JSONFormatter{})

	// 파일 저장
	log.Panic(err.Error())
}
