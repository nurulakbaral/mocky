package logger

import "log"

func Logger (message string, err error) {
	log.Println("⚠️ ", message, err)
}