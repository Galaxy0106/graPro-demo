package main

import (
	"graPro-demo/consumer"
	"graPro-demo/producer"
)

func main() {
	client := producer.StartClient()
	task, _ := producer.NewTask("k-means", " ", 50000, 10000, 1)
	producer.EnQueue(client, task)
	producer.CloseClient(client)
	consumer.StartServer()
	// consumer.CloseServer(server)
}
