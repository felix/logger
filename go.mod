module src.userspace.com.au/felix/logger

go 1.12

require (
	github.com/sirupsen/logrus v1.4.2
	github.com/streadway/amqp v0.0.0-20171101222333-ff791c2d22d3
	github.com/stretchr/testify v1.3.0 // indirect
)

replace src.userspace.com.au/felix/logger => ./
