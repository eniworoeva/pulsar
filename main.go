package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	if err != nil {
		fmt.Println("Error creating client")
	}
	

	defer client.Close()
	
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})
	if err != nil {
		fmt.Println("Error creating client")
	}
	
	
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("Hello World"),
	})
	
	defer producer.Close()
	
	if err != nil {
		fmt.Println("Failed to publish message", err)
	} else {
		fmt.Println("Published message")
	}
}