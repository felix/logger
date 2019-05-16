# Simple structured logger for Go

A simple logger package that provides levels, a number of output formats, and
named sub-logs.  Output formats include key/value, JSON and AMQP/RabbitMQ

## Installation

Install using `go get github.com/felix/logger`.

Documentation is available at http://godoc.org/github.com/felix/logger

## Usage

### Create a key/value logger

```go
log := logger.New(logger.Name("app"), logger.Level(logger.DEBUG))
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
BenchmarkCoreLogger-12           5000000               288 ns/op
BenchmarkLocal-12                2000000               654 ns/op
BenchmarkLogrus-12               1000000              1738 ns/op
BenchmarkFieldsLocal-12          1000000              1024 ns/op
BenchmarkFieldsLogrus-12         1000000              2061 ns/op
```
