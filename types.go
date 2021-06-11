package logpet

import "github.com/sirupsen/logrus"

// StandardLogger is a new type useful to add new methods for default log formats.
type StandardLogger struct {
	*logrus.Logger
	CustomFields  map[string]interface{}
	LogChan       chan Log
	ddAPIKey      string
	ddEndpoint    string
	sendDebugLogs bool
	localMode     bool
}

// Log is a type containing log message and level
type Log struct {
	Message string
	Level   logrus.Level
}
