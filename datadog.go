package logpet

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func (l *StandardLogger) SetupDataDogLogger(datadogEndpoint, datadogAPIKey string, sendDebugLogs bool) error {

	// if provided endpoint is empty we fallback to the default one
	if datadogEndpoint == "" {
		datadogEndpoint = DataDogDefaultEndpoint
	}

	if datadogAPIKey == "" {
		return fmt.Errorf("no API Key provided")
	}

	l.initChannel()
	l.SetDebugMode(sendDebugLogs)

	l.ddAPIKey = datadogAPIKey
	l.ddEndpoint = datadogEndpoint

	// starting log routine
	go l.startLogRoutineListener()

	return nil
}

func (l *StandardLogger) initChannel() {
	l.LogChan = make(chan Log)
}

func (l *StandardLogger) SetDebugMode(debug bool) {
	l.sendDebugLogs = debug
}

func (l *StandardLogger) SendInfoLog(message string) {
	l.LogChan <- Log{
		Message: message,
		Level:   logrus.InfoLevel,
	}
}

func (l *StandardLogger) SendWarnLog(message string) {
	l.LogChan <- Log{
		Message: message,
		Level:   logrus.WarnLevel,
	}
}

func (l *StandardLogger) SendErrLog(message string) {
	l.LogChan <- Log{
		Message: message,
		Level:   logrus.ErrorLevel,
	}
}

func (l *StandardLogger) SendDebugLog(message string) {
	l.LogChan <- Log{
		Message: message,
		Level:   logrus.DebugLevel,
	}
}

func (l *StandardLogger) SendFatalLog(message string) {
	l.LogChan <- Log{
		Message: message,
		Level:   logrus.FatalLevel,
	}
}

func (l *StandardLogger) startLogRoutineListener() {
	var logWriter io.Writer
	var httpClient http.Client
	l.SetOutput(logWriter)

	for logElem := range l.LogChan {

		if !l.sendDebugLogs && logElem.Level == logrus.DebugLevel {
			continue
		}

		newLog := l.AddCustomFields()
		newLog.Message = logElem.Message
		newLog.Level = logElem.Level
		newLog.Time = time.Now()

		err := l.sendLogToDD(newLog, &httpClient)
		if err != nil {
			log.Printf("unable to send log to DataDog, %v", err)
			continue
		}

		if l.sendDebugLogs {
			logBytes, err := newLog.Bytes()
			if err != nil {
				l.SendWarnLog(fmt.Sprintf("error converting log to bytes %v", err))
				continue
			}

			fmt.Println(string(logBytes))
		}

		if logElem.Level == logrus.FatalLevel {
			os.Exit(1)
		}
	}
}

func (l *StandardLogger) sendLogToDD(log *logrus.Entry, httpClient *http.Client) error {

	// obtaining byte slice from log
	logBytes, err := log.Bytes()
	if err != nil {
		return err
	}

	// creating the reader from slice
	body := bytes.NewReader(logBytes)

	// parsing datadog endpoint URL
	urlPrsd, err := url.Parse(l.ddEndpoint)
	if err != nil {
		return err
	}

	// creating new request
	req, err := http.NewRequest(http.MethodPost, urlPrsd.String(), body)
	if err != nil {
		return err
	}

	// adding apikey and content type header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("DD-API-KEY", l.ddAPIKey)

	// doing the request
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// if not ok return an error
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("error when sending logs to DD | Status: %s %v", resp.Status, err))
	}

	return nil
}
