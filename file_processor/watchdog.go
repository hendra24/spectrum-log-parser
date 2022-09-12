package file_processor

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var Sampah = []string{"Comm", "Comms", "IFS"}

func DeleteCollection(ctx context.Context, db *mongo.Database) {

	for i, _ := range Sampah {
		db.Collection(Sampah[i]).Drop(ctx)
	}
}
