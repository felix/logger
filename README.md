# Simple structured logger for Go

[![Build Status](https://travis-ci.org/felix/logger.svg?branch=master)](https://travis-ci.org/felix/logger)

A simple logger package that provides levels, a number of output formats, and
named sub-logs.  Output formats include plain text, key/value, JSON and
AMQP/RabbitMQ

## Installation

Install using `go get github.com/felix/logger`.

Documentation is available at http://godoc.org/github.com/felix/logger

## Usage

### Create a normal logger

```go
log := logger.New(&logger.Options{
	Name:  "app",
	Level: logger.Debug,
})
log.Error("unable to do anything")
```

```text
... [INFO ] app: unable to do anything
```

### Create a key/value logger

```go
import "github.com/felix/logger/outputs/keyvalue"

log := logger.New(&logger.Options{
	Name:      "app",
	Level:     logger.Debug,
    Formatter: keyvalue.New(),
})
log.Warn("invalid something", "id", 344, "error", "generally broken")
```

```text
... [WARN ] app: invalid something id=344 error="generally broken"
```

```text
... [WARN ] app: invalid something id=344 error="generally broken"
```

### Create a sub-logger

```go
sublog := log.Named("database")
sublog.Info("connection initialised")
```

```text
... [INFO ] app.database: connection initialised
```

### Create a new Logger with pre-defined values

For major sub-systems there is no need to repeat values for each log call:

```go
reqID := "555"
msgLog := sublog.WithFields("request", reqID)
msgLog.Error("failed to process message")
```

```text
... [INFO ] app.database: failed to process message request=555
```

## Credits

Solidly based on all the other loggers around, particularly Hashicorp's simple
hclog with additions and modifications as required.
