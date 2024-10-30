package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/FromhelStudio/fromhel-clients-tables-fix/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Migrate(conn *mongo.Client) {
	coll := conn.Database("fromhelClientsDB").Collection("clients")

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			panic(err)
		}

		newId := services.GenerateId()

		date, err := services.ParseDate(doc["registeredAt"])
		if err != nil {
			panic(err)
		}

		name, err := services.Trim(doc["clientName"])
		if err != nil {
			panic(err)
		}

		newDoc := bson.M{
			"_id":          newId,
			"clientId":     doc["clientId"],
			"clientName":   name,
			"updatedAt":    time.Now().Unix(),
			"registeredAt": date,
			"__v":          1,
		}

		_, err = coll.InsertOne(context.TODO(), newDoc)
		if err != nil {
			panic(err)
		}

		_, err = coll.DeleteOne(context.TODO(), bson.M{"_id": doc["_id"]})
		if err != nil {
			panic(err)
		}

		fmt.Println("Migrado -> ", newDoc["_id"])
		fmt.Println("\tNome -> ", newDoc["clientName"])
		fmt.Println("\tVersão -> ", doc["__v"], " -> ", newDoc["__v"])
		fmt.Println("\n\n")
	}
}
