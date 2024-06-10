package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	ogrenciCollection := client.Database("testdb").Collection("ogrenci")

	updateResult, err := ogrenciCollection.UpdateOne(
		context.TODO(),
		bson.M{"no": 21253804},
		bson.D{
			{"$set", bson.D{
				{"isim", "Berkay"},
				{"soyisim", "Aykose"},
				{"sinif", "Üniversite 3"},
			}},
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d satır düzenlendi\n", updateResult.ModifiedCount)
}
