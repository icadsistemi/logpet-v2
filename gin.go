package logpet

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (l *StandardLogger) SendGinRequestLog(ctx *gin.Context, body interface{}) {
	customFields := make(map[string]interface{})

	res, err := json.Marshal(&body)
	if err != nil {
		l.SendWarnfLog("Error during json marshalling in SendGinRequestLog: %v", nil, err)
	}

	customFields["http.body"] = string(res)

	// Remove Bearer token from headers before logging
	headers := make(map[string]interface{}, len(ctx.Request.Header))

	for key, value := range ctx.Request.Header {
		if key == "Authorization" {
			continue
		}
		headers[key] = value[len(value)-1]
	}

	customFields["http.headers"] = headers
	customFields["http.query_parameters"] = ctx.Request.URL.RawQuery
	customFields["http.path_parameters"] = ctx.Request.URL.Path

	l.SendInfofLog("New request received", customFields)
}
