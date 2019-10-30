package logger

import (
	"src.userspace.com.au/logger/message"
	"src.userspace.com.au/logger/writers/json"
	"src.userspace.com.au/logger/writers/kv"
)

func ExampleLevel() {
	keyValue, _ := kv.New(kv.SetTimeFormat(""))
	log, _ := New(
		Name("app"),
		Level(message.DEBUG),
		Writer(keyValue),
	)
	log.Error("unable to do anything")
	// Output: [error] app: unable to do anything
}

func Example_structure() {
	keyValue, _ := kv.New(kv.SetTimeFormat(""))
	log, _ := New(Writer(keyValue))
	log.Warn("invalid something", "id", 344, "error", "generally broken")
	// Output: [warn] invalid something id=344 error="generally broken"
}

func ExampleNamed() {
	keyValue, _ := kv.New(kv.SetTimeFormat(""))
	log, _ := New(
		Name("database"),
		Writer(keyValue),
	)
	log.Error("connection initialised")
	// Output: [error] database: connection initialised
}

func ExampleField() {
	jsonWriter, _ := json.New(json.SetTimeFormat(""))
	log, _ := New(Name("app.database"), Writer(jsonWriter))
	// Create a new Logger with pre-defined values
	reqID := "555"
	msgLog := log.Field("request", reqID)
	msgLog.Error("failed to process message")
	// Output: {"_level":"error","_message":"failed to process message","_name":"app.database","_time":"","request":"555"}
}

func Example_nolevel() {
	keyValue, _ := kv.New(kv.SetTimeFormat(""))
	log, _ := New(Writer(keyValue))
	large := 12345678
	log.Log("metrics or whatnot", "something", large)
	// Output: metrics or whatnot something=12345678
}
