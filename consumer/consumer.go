package consumer

import (
	"be-test/consumer/repository"
	"be-test/consumer/service"
	"be-test/lib"
	"os"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	Consumer sarama.Consumer
}

func (c *KafkaConsumer) Consume(topics []string, signals chan os.Signal) {
	chanMessage := make(chan *sarama.ConsumerMessage, 256)

	for _, topic := range topics {
		partitionList, err := c.Consumer.Partitions(topic)
		if err != nil {
			logrus.Errorf("Unable to get partition: %v", err)
			continue
		}
		for _, partition := range partitionList {
			go consumeMessage(c.Consumer, topic, partition, chanMessage)
		}
	}
	logrus.Infof("Kafka is consuming....")

	var db = lib.DatabaseConnection()
	var balanceRepo = repository.NewBalanceRepository(db)
	var balanceService = service.NewBalanceService(balanceRepo)

ConsumerLoop:
	for {
		select {
		case msg := <-chanMessage:
			if lib.BalancePost.ToKey() == string(msg.Key) {
				err := balanceService.UpdateBalance(msg.Value)
				if err != nil {
					logrus.Errorf("Error update balance consumer: %v", err)
					return
				}
			}
		case sig := <-signals:
			if sig == os.Interrupt {
				break ConsumerLoop
			}
		}
	}
}

func consumeMessage(consumer sarama.Consumer, topic string, partition int32, c chan *sarama.ConsumerMessage) {
	msg, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		logrus.Errorf("Unable to consume partition %v: %v", partition, err)
		return
	}

	defer func() {
		if err := msg.Close(); err != nil {
			logrus.Errorf("Unable to close partition %v: %v", partition, err)
		}
	}()

	for {
		msg := <-msg.Messages()
		c <- msg
	}

}
