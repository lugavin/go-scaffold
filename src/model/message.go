package model

import (
	"time"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	messageCollection = "message"
)

type Message struct {
	Id          string    `bson:"_id",json:"id"`
	Status      string    `json:"status"`
	DeadLetter  bool      `json:"deadLetter"`
	RetryTimes  int       `json:"retryTimes"`
	Exchange    string    `json:"exchange"`
	RoutingKey  string    `json:"routingKey"`
	Body        string    `json:"body"`
	CallbackUrl string    `json:"callbackUrl"`
	Remark      string    `json:"remark"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UpdatedBy   string    `json:"updatedBy"`
}

func InsertMessage(message *Message) error {
	coll := database.Collection(messageCollection)
	message.Id = primitive.NewObjectID().Hex()
	_, err := coll.InsertOne(context.TODO(), message)
	return err
}

func UpdateMessage(id string, message *Message) error {
	coll := database.Collection(messageCollection)
	message.Id = id
	_, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": id}, message)
	return err
}

func GetMessage(id string) (Message, error) {
	coll := database.Collection(messageCollection)
	result := coll.FindOne(context.TODO(), bson.M{"_id": id})
	var message Message
	return message, result.Decode(&message)
}

func DeleteMessage(id string) error {
	coll := database.Collection(messageCollection)
	_, err := coll.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
