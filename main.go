package main

import (
	"be-test/cmd/consumer"
	"be-test/cmd/publisher"
)

func main() {
	go consumer.RunConsumer()
	publisher.RubPublisher()
}
