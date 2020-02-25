package main

import (
	"github.com/nxsre/lumberjack"
	"log"
)

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/example.log",
		TimeFormat: "2006-01-02T15-04",
		MaxSize:    -1,
		MaxAge:     0,
		MaxBackups: 0,
		LocalTime:  false,
		Compress:   true,
	})
	log.Println("Hello World!")
}
