# Simple structured logger for Go

A simple logger package that provides levels, a number of output formats, and
named sub-logs.  Output formats include key/value, JSON and AMQP/RabbitMQ

## Installation

Install using `go get src.userspace.com.au/felix/logger`.

Documentation is available at http://godoc.org/src.userspace.com.au/felix/logger

## Usage

### Create a key/value logger

```go
log := logger.New(logger.SetName("app"), logger.SetLevel(logger.DEBUG))
log.Error("unable to do anything")
```

```text
... [info] app: unable to do anything
```

```go
log.Warn("invalid something", "id", 344, "error", "generally broken")
```

### Add structure

```text
... [warn] app: invalid something id=344 error="generally broken"
```

### Create a sub-logger

```go
sublog := log.GetNamed("database")
sublog.Info("connection initialised")
```

```text
... [info] app.database: connection initialised
```

### Create a new Logger with pre-defined values

For major sub-systems there is no need to repeat values for each log call:

```go
reqID := "555"
msgLog := sublog.WithField("request", reqID)
msgLog.Error("failed to process message")
```

```text
... [info] app.database: failed to process message request=555
```

## Comparison

```
goos: darwin
goarch: amd64
pkg: src.userspace.com.au/felix/logger
BenchmarkLocal-12                2000000               660 ns/op
BenchmarkStdlib-12               5000000               289 ns/op
BenchmarkLogrus-12               1000000              1744 ns/op
BenchmarkFieldsLocal-12          1000000              1029 ns/op
BenchmarkFieldsLogrus-12         1000000              2058 ns/op
```
