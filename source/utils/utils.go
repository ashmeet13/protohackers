package utils

import "github.com/sirupsen/logrus"

var logger *logrus.Entry

func GetLogger() *logrus.Entry {
	if logger != nil {
		return logger
	}
	logger = logrus.NewEntry(logrus.StandardLogger())
	return logger
}
