package logger

import (
	"fmt"

	"github.com/krls08/example-go-submodule/logger/hi"
)

type Logger interface {
	Log(msg string)
}

type DefLogger struct {
}

func NewLogger() Logger {
	return &DefLogger{}
}

func (l *DefLogger) Log(msg string) {
	fmt.Printf("\033[01;33m[LOG]\033[0m %s\n", msg)
	hi.Hello()
}

func main() {
	l := NewLogger()
	l.Log("string")
}
