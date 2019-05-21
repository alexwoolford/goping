package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sparrc/go-ping"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	go func() {}()
	for {
		select {
		case <-ticker.C:
			publishPing()
		}
	}
}

func pingGoogle() (statsJson []byte) {

	pinger, errPinger := ping.NewPinger("www.google.com")
	if errPinger != nil {
		panic(errPinger)
	}

	pinger.Count = 3
	pinger.Run() // blocks until finished
	stats := pinger.Statistics()

	statsJson, err := json.Marshal(stats)
	if err != nil {
		fmt.Println(err)
		return
	}
	return

}

func publishPing() {

	pingGoogleResult := pingGoogle()

	producer, producerErr := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "cp01.woolford.io:9092,cp02.woolford.io:9092,cp03.woolford.io:9092"})
	if producerErr != nil {
		panic(producerErr)
	}

	defer producer.Close()

	topic := "ping"
	_ = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          pingGoogleResult,
	}, nil)

	producer.Flush(1000)

}
