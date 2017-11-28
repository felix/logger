package amqp

import (
	"encoding/json"
	"github.com/felix/logger"
	"github.com/streadway/amqp"
	"io"
	"strings"
	"time"
)

// Writer implementation
type Writer struct {
	url           string
	exchangeName  string
	routingFormat string
	contentType   string
	passive       bool
	conn          *amqp.Connection
	channel       *amqp.Channel
}

// WriterOpt are options for the client
type WriterOpt func(*Writer) error

// New creates a new writer
func New(url, exchange string, opts ...WriterOpt) (*Writer, error) {
	var err error

	w := &Writer{
		url:           url,
		exchangeName:  exchange,
		passive:       false,
		routingFormat: "{name}.{level}",
		contentType:   "application/json",
	}

	// Set variadic options passed
	for _, option := range opts {
		err = option(w)
		if err != nil {
			return nil, err
		}
	}

	w.conn, err = amqp.Dial(w.url)
	if err != nil {
		return nil, err
	}

	w.channel, err = w.conn.Channel()
	if err != nil {
		return nil, err
	}

	if w.passive {
		err = w.channel.ExchangeDeclarePassive(
			w.exchangeName, // name
			"topic",        // type
			true,           // durable
			false,          // auto-deleted
			false,          // internal
			false,          // no-wait
			nil,            // arguments
		)
	} else {
		err = w.channel.ExchangeDeclare(
			w.exchangeName, // name
			"topic",        // type
			true,           // durable
			false,          // auto-deleted
			false,          // internal
			false,          // no-wait
			nil,            // arguments
		)
	}
	if err != nil {
		return nil, err
	}

	return w, nil
}

// SetContentType sets the content type of messages. Default is "application/json"
func SetContentType(ct string) WriterOpt {
	return func(w *Writer) error {
		w.contentType = ct
		return nil
	}
}

// SetPassive makes logger not create exchanges
func SetPassive(p bool) WriterOpt {
	return func(w *Writer) error {
		w.passive = p
		return nil
	}
}

// SetRoutingFormat sets the format for the routing key. Default is "{name}.{level}"
func SetRoutingFormat(k string) WriterOpt {
	return func(w *Writer) error {
		w.routingFormat = k
		return nil
	}
}

// Write implements the logger.MessageWriter interface
func (w Writer) Write(lw io.Writer, m logger.Message) {
	vals := map[string]interface{}{
		"@name":  m.Name,
		"@level": m.Level.String(),
		"@time":  m.Time,
	}

	for i := 0; i < len(m.Fields); i = i + 2 {
		vals[m.Fields[i].(string)] = m.Fields[i+1]
	}

	d, err := json.Marshal(vals)
	if err != nil {
		panic(err)
	}

	msg := amqp.Publishing{
		DeliveryMode: amqp.Transient,
		Timestamp:    time.Now(),
		ContentType:  w.contentType,
		Body:         d,
	}

	routingKey := strings.Replace(w.routingFormat, "{name}", m.Name, 0)
	routingKey = strings.Replace(w.routingFormat, "{level}", m.Level.String(), 0)

	err = w.channel.Publish(
		w.exchangeName,
		routingKey,
		false, // mandatory
		false, // immediate
		msg,
	)
	if err != nil {
		panic(err)
	}
}

// Close implements io.Closer interface
func (w *Writer) Close() (err error) {
	return w.conn.Close()
}