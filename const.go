package logpet

// ====================================
//
// Info Logs Messages
//
//====================================
const (
	httpServerStatusStarted = "Started, listening on port: %s"
)

// ====================================
//
// Warning Logs Messages
//
//====================================
const (
	serverStatusUnauthorized = "%s tried to connect on resource: %s, but it’s unauthorized."
	clientStatusUnauthorized = "I’m trying to connect to %s but I’m unauthorized."
	serverBodyInvalid = "%s sent a request but I don’t know how to read the body."
	clientBodyInvalid = "I sent %s to %s but he can’t read it"
	serverSendResponse = "Can't send the response, error: %s"
	databaseAddError = "Can't add %s for %s"
	databaseGetError = "Can't get %s for %s"
)

// ====================================
//
// Error Logs Messages
//
//====================================
const (

)

// ====================================
//
// Fatal Logs Messages
//
//====================================
const (
	httpServerStartingError = "Error starting HTTP server on port: %s"
	databaseConnectionError = "Invalid connection to the database: %s"
	missingEnvVar = "Missing environment variable: %s"
	missingEntity = "Can't get: %s"
)