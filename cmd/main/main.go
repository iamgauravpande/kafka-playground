package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/iamgauravpande/kafka-playground/pkg/admin"
	"github.com/iamgauravpande/kafka-playground/pkg/producer"
)

func main() {
	brokers := []string{"172.30.217.71:31092"} // defined the broker IP and port : here IP is of WorkderNode and port is NodePort type
	topic := "chat-room"
	admin := admin.NewAdmin(brokers) //create new admin client object
	defer admin.Close()              // close connection at end when all surrounding functions have run
	if !admin.TopicExist(topic) {    // topic Exist method check for topic should not exist.
		admin.TopicCreate(topic) // Create Topic
	}
	username := ""
	fmt.Print("Hi Client , Enter your username: ")
	fmt.Scanln(&username)
	producer := producer.NewProducer(brokers, topic)
	defer producer.Close()
	reader := bufio.NewReader(os.Stdin) // read from console
	for {
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		producer.SendMessage(username, message)
		fmt.Println("Message Sent. Press Ctrl+C to exit")
	}

}
