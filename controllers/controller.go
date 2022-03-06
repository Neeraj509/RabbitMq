package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	M "example.com/m/models"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "GoApp"
const colName = "user"

var client *mongo.Client
var collection *mongo.Collection

func init() {

	clientOption := options.Client().ApplyURI(connectionString)

	client, _ = mongo.Connect(context.TODO(), clientOption)

	fmt.Println("MongoDB connection success")

	collection := client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready", collection)
}

func Input(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user M.Message

	json.NewDecoder(r.Body).Decode(&user)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database(dbName).Collection(colName)

	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
	rabitmq()

}

func rabitmq() {
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

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Published Message to Queue")

}
