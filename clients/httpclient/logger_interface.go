package httpclient

type Logger interface {
	Infof(msg string, args ...interface{})
}
