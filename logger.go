package logpet

import (
	"github.com/sirupsen/logrus"
)

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {

	var standardLogger = &StandardLogger{
		Logger:       logrus.New(),
		CustomFields: make(map[string]interface{}),
	}

	standardLogger.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "date",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyLevel: "status",
		},
	}

	standardLogger.SetReportCaller(true)

	return standardLogger
}

// AddCustomFields adds all the fields in the map CustomFields to our log entry.
// You can call it when you use .Error / .Fatal or other logrus' methods that have Entry as input.
func (l *StandardLogger) AddCustomFields() *logrus.Entry {
	return l.WithFields(l.CustomFields)
}

// ChangeFieldKeys changes the Field keys but if Level, Message or Time are not specified, it uses the defaults.
func (l *StandardLogger) ChangeFieldKeys(fieldMap logrus.FieldMap) {

	if fieldMap[logrus.FieldKeyLevel] == "" {
		fieldMap[logrus.FieldKeyLevel] = "status"
	}

	if fieldMap[logrus.FieldKeyMsg] == "" {
		fieldMap[logrus.FieldKeyMsg] = "message"
	}

	if fieldMap[logrus.FieldKeyTime] == "" {
		fieldMap[logrus.FieldKeyTime] = "date"
	}

	l.Formatter = &logrus.JSONFormatter{
		FieldMap: fieldMap,
	}
}
