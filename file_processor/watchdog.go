package file_processor

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var Sampah = []string{"Comm", "Comms", "IFS"}
var ctx = context.Background()

func DeleteCollection(db *mongo.Database) {

	for i, _ := range Sampah {
		db.Collection(Sampah[i]).Drop(ctx)
	}
}

/*
func CheckFileExist(path string) (bool, error) {
	files, err := ioutil.ReadDir("/tmp/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}
*/
