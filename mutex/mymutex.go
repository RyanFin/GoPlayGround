package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

var mu sync.Mutex

func KafkaProducer(brokerList []string, topic string, msgs []string) {
	producer, err := sarama.NewSyncProducer(brokerList, nil)
	if err != nil {
		log.Fatal("failed to start Kafka producer: ", err)
	}

	defer producer.Close()

	mu.Lock()
	defer mu.Unlock()

	// this content is locked within the Mutex
	for _, msg := range msgs {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(msg),
		}

		_, _, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("failed to send msg: ", err)
		} else {
			fmt.Printf("Message sent: %s\n", msg)
		}
	}

	time.Sleep(2 * time.Second)
}

func KafkaConsumer(brokerList []string, topic string) {
	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		log.Fatal("failed to start kafka consumer: ", err)
	}

	defer consumer.Close()

	// subscribe to the Kafka topic
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("failed to start consuming partition: ", err)
	}

	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		mu.Lock()
		fmt.Println("received msg", string(msg.Value))
		time.Sleep(1 * time.Second)
		mu.Unlock()
	}
}

func main() {
	brokerList := []string{"localhost:9092"} // Kafka broker address
	topic := "test-topic"
	messages := []string{
		"Message 1",
		"Message 2",
		"Message 3",
	}

	go KafkaProducer(brokerList, topic, messages)

	go KafkaConsumer(brokerList, topic)

	time.Sleep(10 * time.Second)
}
