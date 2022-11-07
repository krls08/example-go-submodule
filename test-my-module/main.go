package main

import (
	"github.com/krls08/example-go-submodule/logger"
	//"github.com/krls08/example-go-submodule/logger/hi"
	"github.com/krls08/go_logger_ex"
)

func main() {
	go_logger_ex.Log("hi!")
	l := logger.NewLogger()
	l.Log("test")

	//hi.Hello()

}
