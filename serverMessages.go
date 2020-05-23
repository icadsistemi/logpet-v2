package logpet

// ====================================
//
// Info Logs Messages
//
//====================================

// HTTPServerStarted prints the string "Started, listening on port: %s" with an Info error.
// It accepts a string as port where the server listens.
func (l *StandardLogger) HTTPServerStarted(port string) {
	l.AddCustomFields().Infof(httpServerStatusStarted, port)
}

// ====================================
//
// Warning Logs Messages
//
//====================================

// HTTPServerUnauthorizedResponse prints the string "%s tried to connect on resource: %s, but it’s unauthorized." with a Warning error.
// It accepts two strings representing the host that made the request and the resource that it wants.
func (l *StandardLogger) HTTPServerUnauthorizedResponse(request, resource  string) {
	l.AddCustomFields().Warningf(serverStatusUnauthorized, request, resource)
}

// HTTPServerInvalidBodyResponse prints the string "%s sent a request but I don’t know how to read the body." with a Warning error.
// It accepts a string representing the host that made the request.
func (l *StandardLogger) HTTPServerInvalidBodyResponse(request string) {
	l.AddCustomFields().Warningf(serverBodyInvalid, request)
}

// HTTPServerSendResponseError prints the string "Can't send the response, error: %s" with a Warning error.
// It accepts a string representing the error
func (l *StandardLogger) HTTPServerSendResponseError(error string) {
	l.AddCustomFields().Warningf(serverSendResponse, error)
}


// ====================================
//
// Fatal Logs Messages
//
//====================================

// InvalidDatabaseConnection prints the string "Invalid connection to the database: %s" with a Fatal error.
// It accepts a string that could be the hostname of the database.
func (l *StandardLogger) HTTPServerStartingError(port string) {
	l.AddCustomFields().Fatalf(httpServerStartingError, port)
}