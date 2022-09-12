package db_connector

import (
	"context"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"grade"`
}

type Kinerja struct {
	Time time.Time `bson:"time"`
}

type Spectrum struct {
	Time     time.Time `bson:"time"`
	B1       string    `bson:"b1"`
	B2       string    `bson:"b2"`
	B3       string    `bson:"b3"`
	Elem     string    `bson:"elem"`
	VarText  string    `bson:"vartext"`
	Status   string    `bson:"status"`
	Tag      string    `bson:"tag"`
	Operator string    `bson:"operator"`
	MeCl     string    `bson:"mecl"`
}

func Connect(ctx context.Context, db_name string) (*mongo.Database, error) {

	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	//return the database that we uses
	log.Println("Successfull connect to db " + db_name)
	return client.Database(db_name), nil
}

func InsertToDB(ctx context.Context, db *mongo.Database, data []string) {

	_, val := checkCollectionExist(ctx, db, data[2])

	if val {
		t, err1 := ConvertTimeToFormat(data[0])
		if err1 != nil {
			return
		} else {
			_, err := db.Collection(data[2]).InsertOne(ctx, Spectrum{t, data[2], data[3], data[4], data[5], data[6], data[7], data[8], data[9], data[10]})
			if err != nil {
				log.Fatal(err)
			}
		}

	} else {
		db.CreateCollection(ctx, data[2])
		t, err1 := ConvertTimeToFormat(data[0])
		if err1 != nil {
			return
		} else {
			_, err := db.Collection(data[2]).InsertOne(ctx, Spectrum{t, data[2], data[3], data[4], data[5], data[6], data[7], data[8], data[9], data[10]})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func checkCollectionExist(ctx context.Context, db *mongo.Database, str string) (error, bool) {

	names, err := db.ListCollectionNames(ctx, bson.D{{"name", str}})

	if err != nil {
		return err, false
	}

	if names != nil {
		return nil, true
	}

	return nil, false
}

func ConvertTimeToFormat(str string) (time.Time, error) {
	datas := strings.Split(str, " ")
	valid := datas[0] + " " + datas[1][0:8]

	t, err := time.Parse("02.01.2006 15:04:05", valid)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

/*
func FindByName(str string) {
	d, err := Connect()
	if err != nil {
		log.Fatal(err.rror()
	}

csr, err := db.Collection(str).Find(ctx, bson.M{"name": str})
if err != nil {
		log.Fatal(err.rror())
	}

defer csr.Close(ctx)

result := make([]stuent, 0)
for csr.Next(ctx) {
		var row student
		err := csr.Decoe(&ow)
		if err != nil {
			log.Fatal(err.rror())
		}
		rsult = append(result,row)
	}

if len(result) > 0 {
	for i := range resut {
			fmt.Println("inde :",i)
			fmt.Println("name	:",reslt[i].Name)
			fmt.Println("grade	:", rsult[i].Grad)
		}

} ese {
	fmt.Pritln("empty entries")
	}
}

uc removeByName(str string) {
db, err := Connect()
	if err != nil {
		log.Fatal(err.rror()
	}

va selector = bson.M{"name": str}
_, err = db.Collection("student").eleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.rror())
	}
	ft.Println("remove sucesfully")
}
*/
