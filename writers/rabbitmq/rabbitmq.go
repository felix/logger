package amqp

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/streadway/amqp"
	"src.userspace.com.au/logger/internal"
	"src.userspace.com.au/logger/message"
)

// Writer implementation
type Writer struct {
	url           string
	exchangeName  string
	routingFormat string
	exchangeType  string
	contentType   string
	passive       bool
	conn          *amqp.Connection
	channel       *amqp.Channel
}

// New creates a new writer
func New(url, exchange string, opts ...WriterOpt) (*Writer, error) {
	var err error

	w := &Writer{
		url:           url,
		exchangeName:  exchange,
		passive:       false,
		exchangeType:  "topic",
		routingFormat: "{name}.{level}",
		contentType:   "application/json",
	}

	// Set variadic options passed
	for _, option := range opts {
		if err = option(w); err != nil {
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
			w.exchangeType, // type
			true,           // durable
			false,          // auto-deleted
			false,          // internal
			false,          // no-wait
			nil,            // arguments
		)
	} else {
		err = w.channel.ExchangeDeclare(
			w.exchangeName, // name
			w.exchangeType, // type
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

// Write implements the logger.Writer interface
func (w Writer) Write(m message.Message) {
	vals := map[string]interface{}{
		"@name":    m.Name,
		"@level":   m.Level.String(),
		"@message": m.Content,
		"@time":    m.Time,
	}

	for k, v := range m.Fields {
		vals[k] = v
	}

	if len(m.Extras) > 0 {
		// Allow for an odd number of extras
		offset := len(m.Extras) % 2
		if offset != 0 {
			for k, v := range m.Extras {
				vals[fmt.Sprintf("extra%02d", k)] = v
			}
		} else {
			for i := offset; i < len(m.Extras); i = i + 2 {
				vals[internal.ToString(m.Extras[i])] = m.Extras[i+1]
			}
		}
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

	if err := w.channel.Publish(
		w.exchangeName,
		routingKey,
		false, // mandatory
		false, // immediate
		msg,
	); err != nil {
		panic(err)
	}
}

// Close implements io.Closer interface
func (w *Writer) Close() (err error) {
	return w.conn.Close()
}
