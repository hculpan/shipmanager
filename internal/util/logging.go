package util

import "log"

func LogError(err string) {
	log.Printf("ERROR: %s", err)
}
