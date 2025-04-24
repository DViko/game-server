package helpers

import "log"

func ErrorHelper(err error, message string) {

	if err != nil {
		log.Printf(message, err)
	}
}
