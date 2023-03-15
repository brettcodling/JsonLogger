# JsonLogger [![Go Report Card](https://goreportcard.com/badge/github.com/brettcodling/jsonlogger)](https://goreportcard.com/report/github.com/brettcodling/jsonlogger)

A Go (golang) json logging library for use with GCP Logs Explorer

## Installation

As a library

```shell
go get github.com/brettcodling/jsonlogger
```

## Usage

```go
package main

import github.com/brettcodling/jsonlogger"

func main() {
  // You can use the built in log level methods (Info, Debug, Warn, Error)
  jsonlogger.Info("Info level log")
  
  // You can pass a context interface to the log as long as it can be marshalled to json
  err := ...
  jsonlogger.Error("Error!!!", map[string]string{"error": err.Error()})
  
  // You can also manually populate the JsonLog struct and call Log()
  jsonlogger.JsonLog{
    Message:  "Manual log",
    Context:  map[string]int{"test": 1}
    Severity: "INFO",
    Type:     "log",
  }.Log()
}
```
