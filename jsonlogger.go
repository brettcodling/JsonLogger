package jsonlogger

import (
	"encoding/json"
	"log"
	"os"
)

var (
	logLevels = map[string]int{
		"INFO":  0,
		"DEBUG": 1,
		"WARN":  2,
		"ERROR": 3,
	}
	logLevel = logLevels[os.Getenv("LOG_LEVEL")]
)

type JsonLog struct {
	Message  string      `json:"message"`
	Context  interface{} `json:"context"`
	Type     string      `json:"type"`
	Severity string      `json:"severity"`
}

// log will log a json message
// If the message cannot be marshalled the json error will be logged
func (jLog JsonLog) Log() {
	// Set default parameters
	if jLog.Type == "" {
		jLog.Type = "log"
	}
	if jLog.Severity == "" {
		jLog.Severity = "INFO"
	}
	if level, ok := logLevels[jLog.Severity]; ok {
		if level < logLevel {
			return
		}
	}
	encodedLog, err := json.Marshal(jLog)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(encodedLog))
}

// Info will log a json message with a severity of INFO
func Info(msg string, context ...interface{}) {
	var c interface{}
	switch {
	case len(context) == 1:
		c = context[0]
	case len(context) > 1:
		c = context
	}
	JsonLog{
		Message:  msg,
		Context:  c,
		Severity: "INFO",
		Type:     "log",
	}.Log()
}

// Error will log a json message with a severity of ERROR
func Error(msg string, context ...interface{}) {
	var c interface{}
	switch {
	case len(context) == 1:
		c = context[0]
	case len(context) > 1:
		c = context
	}
	JsonLog{
		Message:  msg,
		Context:  c,
		Severity: "ERROR",
		Type:     "log",
	}.Log()
}

// Warn will log a json message with a severity of WARNING
func Warn(msg string, context ...interface{}) {
	var c interface{}
	switch {
	case len(context) == 1:
		c = context[0]
	case len(context) > 1:
		c = context
	}
	JsonLog{
		Message:  msg,
		Context:  c,
		Severity: "WARNING",
		Type:     "log",
	}.Log()
}

// Debug will log a json message with a severity of DEBUG
func Debug(msg string, context ...interface{}) {
	var c interface{}
	switch {
	case len(context) == 1:
		c = context[0]
	case len(context) > 1:
		c = context
	}
	JsonLog{
		Message:  msg,
		Context:  c,
		Severity: "DEBUG",
		Type:     "log",
	}.Log()
}
