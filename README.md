# Simple structured logger for Go

A simple logger package that provides levels, a number of output formats, and
named sub-logs.  Output formats include key/value, JSON, null and AMQP/RabbitMQ

## Installation

Install using `go get src.userspace.com.au/felix/logger`.

Documentation is available at http://godoc.org/src.userspace.com.au/felix/logger

## Usage

There is a package level logger that is set to level 'WARN'.

### Create a key/value logger

```go
log := logger.New(logger.Name("app"), logger.Level(logger.DEBUG))
log.Error("unable to do anything")
```

```text
... [info] app: unable to do anything
```

### Add structure

```go
log.Warn("invalid something", "id", 344, "error", "generally broken")
```

```text
... [warn] app: invalid something id=344 error="generally broken"
```

### Create a named sub-logger

```go
sublog := log.Named("database")
sublog.Info("connection initialised")
```

```text
... [info] app.database: connection initialised
```

### Create a new Logger with pre-defined values

For major sub-systems there is no need to repeat values for each log call:

```go
reqID := "555"
msgLog := sublog.Field("request", reqID)
msgLog.Error("failed to process message")
```

```text
... [info] app.database: failed to process message request=555
```

There is also a Log command with no defined level. These messages are always
printed:

```go
log.Log("metrics or whatnot", "something", large)
```

```text
... metrics or whatnot something="12345678"
```

## Comparison

```
BenchmarkCoreLogger-12           5000000               288 ns/op
BenchmarkLocal-12                2000000               654 ns/op
BenchmarkLogrus-12               1000000              1738 ns/op
BenchmarkFieldsLocal-12          1000000              1024 ns/op
BenchmarkFieldsLogrus-12         1000000              2061 ns/op
```
