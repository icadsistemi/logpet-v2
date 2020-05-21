package logpet


// ====================================
//
// Info Logs Messages
//
//====================================



// ====================================
//
// Fatal Logs Messages
//
//====================================

// InvalidDatabaseConnection prints the string "Invalid connection to the database: %s" with a Fatal error.
// It accepts a string that could be the hostname of the database.
func (l *StandardLogger) InvalidDatabaseConnection(databaseHost string) {
	l.AddCustomFields().Errorf(databaseConnectionError, databaseHost)
}


