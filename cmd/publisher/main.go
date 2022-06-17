package publisher

import (
	"be-test/producer"
	"net/http"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func RubPublisher() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	kafkaConfig := getKafkaConfig("", "")
	producers, err := sarama.NewSyncProducer([]string{"0.0.0.0:9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer: %v", err)
		return
	}

	defer func() {
		if err := producers.Close(); err != nil {
			logrus.Errorf("Unable to stop kafka producer: %v", err)
			return
		}
	}()

	logrus.Infof("Success create kafka sync-producer")

	logrus.Infof("Running http server")
	logrus.Fatal(http.ListenAndServe(":8080", producer.RestRouter(producers)))
}

func getKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}

	return kafkaConfig
}
