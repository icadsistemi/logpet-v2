package logpet

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func (l *StandardLogger) LogAPIGatewayProxyRequest(event events.APIGatewayProxyRequest) {
	var message = "Received API Gateway Proxy Request."

	if event.Body != "" {
		message = fmt.Sprintf("%s Body: %s - ", message, event.Body)
	}

	email, _ := ReadEmailFromAPIGatewayProxyRequestEvent(event)

	l.CustomFields["http.query_parameters"] = event.QueryStringParameters
	l.CustomFields["http.path_parameters"] = event.PathParameters
	l.CustomFields["http.headers"] = event.Headers
	l.CustomFields["http.method"] = event.HTTPMethod
	l.CustomFields["http.url"] = event.Path
	l.CustomFields["user.name"] = email
	l.AddCustomFields().Debug(message)
}

func getClaimsFromAPIGW(event events.APIGatewayProxyRequest) (map[string]interface{}, error) {
	claimsRaw, ok := event.RequestContext.Authorizer["claims"]
	if !ok {
		return nil, fmt.Errorf("unable to find claims field in event")
	}

	return claimsRaw.(map[string]interface{}), nil
}

// ReadEmailFromAPIGatewayProxyRequestEvent takes email field in cognito token
func ReadEmailFromAPIGatewayProxyRequestEvent(event events.APIGatewayProxyRequest) (string, error) {

	claims, err := getClaimsFromAPIGW(event)
	if err != nil {
		return "", err
	}

	emailRaw, ok := claims["email"]
	if !ok {
		return "", fmt.Errorf("unable to find email field in event")
	}

	return emailRaw.(string), nil
}
