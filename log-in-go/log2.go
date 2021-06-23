package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("info2.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	requestLogger := log.WithFields(log.Fields{"request_id": "request_id", "user_ip": "user_ip"})
	requestLogger.Info("something happened on that request") // will log request_id and user_ip
	requestLogger.Warn("something not great happened")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"number": 1,
		"size":   10,
	}).Info("A walrus appears")
}
