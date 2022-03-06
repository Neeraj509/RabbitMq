package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/streadway/amqp"
)

var count = 0

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
func CreateFile(result int) {

	file, err := os.Create("NewFile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, fmt.Sprintf("%d\n", result))

	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("NewFile.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	checkNilErr(err)

	fmt.Println("Total request count: ", string(databyte))

}

func main() {

	fmt.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	// with this channel open, we can then start to interact
	// with the instance and declare Queues that we can publish and
	// subscribe to
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
			UpdateRequestHit()
		}
	}()
	

	// fmt.Println("Successfully Connected to our RabbitMQ Instance")

	fmt.Println(" [*] - Waiting for messages")
	<-forever
	

}
func UpdateRequestHit() {
	databyte, err := ioutil.ReadFile("NewFile.txt")
	if err != nil {
		CreateFile(1)
	} else {
		
		l := string(databyte)
		l = strings.TrimSuffix(l, "\n")
		i, _ := strconv.Atoi(l)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		CreateFile(i + 1)

	}


}
