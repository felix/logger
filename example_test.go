package logger

import (
	"os"

	"src.userspace.com.au/logger/writers/json"
	"src.userspace.com.au/logger/writers/kv"
)

func ExampleDebug() {
	keyValue, _ := kv.New(kv.SetTimeFormat(""))
	os.Setenv("DEBUG", "true")
	log, _ := New(
		Name("app"),
		Writer(keyValue),
	)
	log.Debug("unable to do anything")
	// Output: app: unable to do anything
}

func Example_structure() {
	keyValue, _ := kv.New(kv.SetTimeFormat(""))
	log, _ := New(Writer(keyValue))
	log.Fields("id", 344, "error", "generally broken").Log("invalid something")
	// Output: invalid something id=344 error="generally broken"
}

func ExampleNamed() {
	keyValue, _ := kv.New(kv.SetTimeFormat(""))
	log, _ := New(
		Name("database"),
		Writer(keyValue),
	)
	log.Info("connection initialised")
	// Output: database: connection initialised
}

func ExampleField() {
	jsonWriter, _ := json.New(json.SetTimeFormat(""))
	log, _ := New(Name("app.database"), Writer(jsonWriter))
	// Create a new Logger with pre-defined values
	reqID := "555"
	msgLog := log.Field("request", reqID)
	msgLog.Log("failed to process message")
	// Output: {"_message":"failed to process message","_name":"app.database","_time":"","request":"555"}
}
