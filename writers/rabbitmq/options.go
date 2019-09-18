package amqp

// WriterOpt are options for the client
type WriterOpt func(*Writer) error

// ContentType sets the content type of messages. Default is "application/json"
func tContentType(ct string) WriterOpt {
	return func(w *Writer) error {
		w.contentType = ct
		return nil
	}
}

// Passive makes logger not create exchanges
func Passive(p bool) WriterOpt {
	return func(w *Writer) error {
		w.passive = p
		return nil
	}
}

// ExchangeType sets the exchange type from the default 'topic'.
func ExchangeType(t string) WriterOpt {
	return func(w *Writer) error {
		w.exchangeType = t
		return nil
	}
}

// RoutingFormat sets the format for the routing key. Default is "{name}.{level}"
func RoutingFormat(k string) WriterOpt {
	return func(w *Writer) error {
		w.routingFormat = k
		return nil
	}
}
