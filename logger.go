package logpet

import "github.com/sirupsen/logrus"

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var standardLogger = &StandardLogger{logrus.New(), make(map[string]interface{})}

	standardLogger.Formatter = &logrus.JSONFormatter{
		FieldMap:    logrus.FieldMap{logrus.FieldKeyTime: "date"},
		PrettyPrint: true,
	}

	standardLogger.SetReportCaller(true)

	return standardLogger
}

// AddCustomFields adds all the fields in the map CustomFields to our log entry.
// You can call it when you use .Error / .Fatal or other logrus' methods that have Entry as input.
func (l *StandardLogger) AddCustomFields() *logrus.Entry {
	return l.WithFields(l.CustomFields)
}