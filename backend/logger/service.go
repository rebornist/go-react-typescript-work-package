package logger

type LoggerService interface {
	PrintLogger(map[string]interface{}, int, string) error
	WriteLogger(map[string]interface{}, int, string) error
	RaiseError(error)
}
