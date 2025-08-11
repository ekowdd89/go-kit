package main

import (
	"context"
	"fmt"
	// "log"
	// "strings"
	kafka "github.com/ekowdd89/go-kit/internals/kafka"
	// kafka "github.com/segmentio/kafka-go"
)


// func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
// 	brokers := strings.Split(kafkaURL, ",")
// 	return kafka.NewReader(kafka.ReaderConfig{
// 		Brokers:  brokers,
// 		GroupID:  groupID,
// 		Topic:    topic,
// 		MinBytes: 10e3, // 10KB
// 		MaxBytes: 10e6, // 10MB
// 	})
// }

func main(){
	kf, err:= kafka.New(
		kafka.WithBrokers([]string{"host.docker.internal:19092"}),
		kafka.WithTopic("author"),
		kafka.WithPubSubName("author"),
	)
	if err != nil {
		panic(err)
	}
	if err := kf.Subscribe(context.Background(), "default-group", func(val []byte) {
		fmt.Println(string(val))
	}); err != nil {
		panic(err)
	}
}

func Reader(){
	// reader := getKafkaReader("host.docker.internal:19092", "author", "author-group")
	// for {
	// 	m, err := reader.ReadMessage(context.Background())
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Println(m.Topic, m.Key, string(m.Value))
	// }
}