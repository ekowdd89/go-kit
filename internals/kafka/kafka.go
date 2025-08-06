package kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)




type OptsFunc func(*Kafka) error


func WithBrokers(brokers []string) OptsFunc {
	return func(k *Kafka) (err error) {
		k.brokers = brokers
		return
	}
}

func WithTopic(topic string) OptsFunc {
	return func(k *Kafka) (err error) {
		k.topic = topic
		return
	}
}
type Kafka struct {
	brokers []string
	topic string
	sender *kafka_sarama.Sender
	client cloudevents.Client
}



func New(opts ...OptsFunc) (kf *Kafka, err error){
	kf = &Kafka{}

	for _, opt:= range opts {
		if err = opt(kf); err !=nil{
			return
		}
	}
	if len(kf.brokers) == 0 {
		return nil, fmt.Errorf("missing brokers")
	}
	saramaConfig:= sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0

	sender, err:= kafka_sarama.NewSender(kf.brokers, saramaConfig, kf.topic)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate cloudevents kafka sender")
	}
	kf.client, err = cloudevents.NewClient(sender, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate cloudevents client")
	}
	return
}

func (kf *Kafka)Close(ctx context.Context) (err error) {
	if kf.sender == nil {
		return nil
	}
	return kf.client.Close()
}