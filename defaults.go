package logpet


// ====================================
//
// Info Logs Messages
//
//====================================


// ====================================
//
// Warning Logs Messages
//
//====================================

// DatabaseAddingError prints the string "Can't add %s for %s" with a Warning error.
// It accepts two strings representing the entity we want to add and the user who requested it.
func (l *StandardLogger) DatabaseAddingError(entity, user string) {
	l.AddCustomFields().Warningf(databaseAddError, entity, user)
}

// DatabaseGetError prints the string "Can't get %s for %s" with a Warning error.
// It accepts two strings representing the entity we want to add and the user who requested it.
func (l *StandardLogger) DatabaseGetError(entity, user string) {
	l.AddCustomFields().Warningf(databaseGetError, entity, user)
}


// ====================================
//
// Fatal Logs Messages
//
//====================================

// InvalidDatabaseConnection prints the string "Invalid connection to the database: %s" with a Fatal error.
// It accepts a string that could be the hostname of the database.
func (l *StandardLogger) InvalidDatabaseConnection(databaseHost string) {
	l.AddCustomFields().Fatalf(databaseConnectionError, databaseHost)
}

// MissingEnvVariable prints the string "Missing environment variable: %s" with a Fatal error.
// It accepts as a string, the name of the environment variable.
func (l *StandardLogger) MissingEnvVariable(env string) {
	l.AddCustomFields().Fatalf(missingEnvVar, env)
}

// MissingNecessaryEntity prints the string "Can't get: %s" with a Fatal error.
// It accepts as a string, the name of the entity name.
func (l *StandardLogger) MissingNecessaryEntity(ent string) {
	l.AddCustomFields().Fatalf(missingEntity, ent)
}

