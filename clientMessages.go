package logpet

// ====================================
//
// Warning Logs Messages
//
//====================================

// HTTPClientUnauthorized prints the string "I’m trying to connect to %s but I’m unauthorized." with a Warning error.
// It accepts a string representing the server to which we made the request.
func (l *StandardLogger) HTTPClientUnauthorized(server  string) {
	l.AddCustomFields().Warningf(clientStatusUnauthorized, server)
}

// HTTPClientInvalidBody prints the string "I sent %s to %s but he can’t read it" with a Warning error.
// It accepts two strings: the body of the request and the server to which we made the request.
func (l *StandardLogger) HTTPClientInvalidBody(body, server string) {
	l.AddCustomFields().Warningf(clientBodyInvalid, body, server)
}
