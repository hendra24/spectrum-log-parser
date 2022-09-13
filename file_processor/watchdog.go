package file_processor

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var Sampah = []string{"Comm", "Comms ", "IFS "}

func DeleteCollection(ctx context.Context, db *mongo.Database) error {

	for i, _ := range Sampah {
		err := db.Collection(Sampah[i]).Drop(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
