package errors

import "log"

func Check(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
