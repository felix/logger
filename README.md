# Simple structured logger for Go

A simple logger package that provides a number of output formats, and
named sub-logs.  Output formats include key/value, JSON, null and AMQP/RabbitMQ

## Installation

Install using `go get src.userspace.com.au/logger`.

Documentation is available at http://godoc.org/src.userspace.com.au/logger

## Usage

There is a package level logger with two levels, debug and not!

### Create a key/value logger

```go
log := logger.New(logger.Name("app"))
log.Log("unable to do anything")
```

```text
... app: unable to do anything
```

### Add structure

```go
log.Log("invalid something", "id", 344, "error", "generally broken")
```

```text
... app: invalid something id=344 error="generally broken"
```

### Create a named sub-logger

```go
sublog := log.Named("database")
sublog.Log("connection initialised")
```

```text
... app.database: connection initialised
```

### Create a new Logger with pre-defined values

For major sub-systems there is no need to repeat values for each log call:

```go
reqID := "555"
msgLog := sublog.Field("request", reqID)
msgLog.Log("failed to process message")
```

```text
... app.database: failed to process message request=555
```

## Comparison

```
BenchmarkCoreLogger-12           4731555               243 ns/op
BenchmarkLocal-12                2035790               597 ns/op
BenchmarkLogrus-12                698662              1725 ns/op
BenchmarkFieldsLocal-12          1000000              1010 ns/op
BenchmarkFieldsLogrus-12          592014              2022 ns/op
```
