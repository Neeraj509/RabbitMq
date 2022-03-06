package model

type Message struct {
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
	File string `json:"file" bson:"file"`
}
