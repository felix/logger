package logger

import (
	"io/ioutil"
	"log"
	"testing"

	logrus "github.com/sirupsen/logrus"
	"src.userspace.com.au/logger/message"
	"src.userspace.com.au/logger/writers/kv"
)

func dummyWriter() message.Writer {
	kv, err := kv.New(kv.SetOutput(ioutil.Discard))
	if err != nil {
		panic("failed to create keyvalue")
	}
	return kv
}

func BenchmarkCoreLogger(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	for n := 0; n < b.N; n++ {
		log.Print("Some text")
	}
}

func BenchmarkLocal(b *testing.B) {
	l, _ := New(Writer(dummyWriter()))
	for n := 0; n < b.N; n++ {
		l.Log("Some text")
	}
}

func BenchmarkLogrus(b *testing.B) {
	logrus.SetOutput(ioutil.Discard)
	for n := 0; n < b.N; n++ {
		logrus.Error("Some text")
	}
}

func BenchmarkFieldsLocal(b *testing.B) {
	l, _ := New(Writer(dummyWriter()))
	l.Field("key", "value")
	l.Field("one", "two")
	for n := 0; n < b.N; n++ {
		l.Info("Some text")
	}
}

func BenchmarkFieldsLogrus(b *testing.B) {
	logrus.SetOutput(ioutil.Discard)
	l := logrus.WithFields(logrus.Fields{"key": "value", "one": "two"})
	for n := 0; n < b.N; n++ {
		l.Error("Some text")
	}
}
