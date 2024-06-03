package generic

import (
	"log"
)

func LogWarningf(format string, v ...any) {
	log.Printf(format, v...)
}

func LogError(e error) {
	if e != nil {
		log.Println(e)
	}
}
