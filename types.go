package logpet

import "github.com/sirupsen/logrus"

// StandardLogger is a new type useful to add new methods for default log formats.
type StandardLogger struct {
	*logrus.Logger
	CustomFields map[string]interface{}
}