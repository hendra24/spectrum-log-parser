package file_processor

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	db_connector "github.com/hendra24/spectrum-log-parser/db"
	"go.mongodb.org/mongo-driver/mongo"
)

const DATA_LOGS_PATH = "H:\\GO\\mybefail\\logs\\20220328\\"
const DATA_WAREHOUSE_PATH = "H:\\GO\\mybefail\\warehouse\\"

/*
func ProcessFile(fname string) error {
	//read file from pool folder
	files, err := ioutil.ReadDir(DATA_LOGS_PATH)

	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		db, err := db_connector.Connect("test_db")
		if err != nil {
			log.Fatal(err)
		}
		readFile(f.Name(), "\t", db)
	}
	return nil
}
*/
func ReadFile(fname string, sep string, db *mongo.Database) error {
	//count execution tme
	start := time.Now()

	//acces log file
	file, err := os.Open(DATA_LOGS_PATH + fname)
	if err != nil {
		log.Fatal("Error when read file :" + err.Error())
		return err
	}
	defer file.Close()
	log.Println("Processing file : ", fname)
	scanner := bufio.NewScanner(file)

	//readfile
	for scanner.Scan() {
		//datas := bytes.Split(scanner.Bytes(), []bte{9})
		//split data by tab to collection of string
		datas := strings.Split(strings.Replace(scanner.Text(), "  ", "", -1), sep)

		if len(datas) == 13 && len(datas[0]) >= 19 && datas[2] != "" {
			//saveToFile(datas, data[2])
			//save file to mongoo db
			db_connector.InsertToDB(db, datas)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return err
	}

	duration := time.Since(start)
	log.Println("Done Processing: ", fname, "in", duration.Seconds(), "second")
	DeleteCollection(db)
	log.Println("DB SAMPAH DI Hapus")

	return nil
}

func saveToFile(s []string, fname string) {
	// If the file doesn't exist,create it, or append to the file
	os.Chdir(DATA_WAREHOUSE_PATH)
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error while save file :" + err.Error())
	}
	//set length data to save
	data := strings.Join(s[1:12], ";")

	if _, err := f.WriteString(data + "\n"); err != nil {
		f.Close() // inore error; Write error takes precedence
		log.Fatal(err)
	}
	defer f.Close()
}
