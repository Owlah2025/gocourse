package main

import "github.com/sirupsen/logrus"

//go mod init gocourse
//go get github.com/sirupsen/logrus

func main() {
	log := logrus.New()

	//set log level
	log.SetLevel(logrus.InfoLevel)

	//set log format
	log.SetFormatter(&logrus.JSONFormatter{}) //JSON is a javascript object notation

	//logging examples
	log.Info("This is an info message.")
	log.Warn("This is a warning message.")
	log.Error("This is an error message.")

	//remember that the value of the key in the map is interface{} wich is any , do you remember!?
	log.WithFields(logrus.Fields{ // remember that map[] dosn't maintain order !!
		"username": "Ahmed Hazem",
		"method":   "GET",
	}).Info("User logged in") //chaining methods
}
