package infrastructure

import "log"

func Logger(args ...interface{}) {
	log.Println(args...)
}
