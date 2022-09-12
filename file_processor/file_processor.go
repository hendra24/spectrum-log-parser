package file_processor

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	db_connector "github.com/hendra24/spectrum-log-parser/db"
	"go.mongodb.org/mongo-driver/mongo"
)

const DATA_LOGS_PATH = "H:\\GO\\mybefail\\logs\\20220328\\"
const DATA_WAREHOUSE_PATH = "H:\\GO\\mybefail\\warehouse\\"

func ReadFile(ctx context.Context, fname string, sep string, db *mongo.Database) error {
	//count execution tme
	start := time.Now()
	//acces log file
	file, err := os.Open(DATA_LOGS_PATH + fname)
	if err != nil {
		return err
	}

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
			db_connector.InsertToDB(ctx, db, datas)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	duration := time.Since(start)
	defer log.Println("Done Processing: ", fname, "in", duration.Seconds(), "second")
	//force reading file to close to avoid infinite
	file.Close()
	err = moveFile(DATA_LOGS_PATH+fname, DATA_WAREHOUSE_PATH+fname)
	if err != nil {
		return err
	}
	//DELETE UNSUEFUL LOG
	err = DeleteCollection(ctx, db)
	if err != nil {
		return err
	} else {
		log.Println("DB SAMPAH DI Hapus")
	}

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
	f.Close()
}

func moveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}
