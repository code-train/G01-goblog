package logger

import "log"

// LogError method
func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}
