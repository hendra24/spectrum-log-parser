package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	queue "github.com/hendra24/spectrum-log-parser/data_structure"
	db_connector "github.com/hendra24/spectrum-log-parser/db"
	"github.com/hendra24/spectrum-log-parser/file_processor"
)

const DATA_LOGS_PATH = "H:\\GO\\mybefail\\logs\\20220328\\"

func main() {

	//initialize new queue
	fileToProcess := queue.NewQueue("Parser File to DB")
	var jobs []queue.Job

	//check directory folder have file ? if y
	for {
		log.Println("Checking log in path " + string(DATA_LOGS_PATH))
		files, err := ioutil.ReadDir(DATA_LOGS_PATH)
		if err != nil {
			log.Fatal(err)
		}

		//connect to db
		db, err := db_connector.Connect("test_db")
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}

			action := func() error {
				err := file_processor.ReadFile(f.Name(), "\t", db)
				if err != nil {
					return err
				}
				return nil
			}
			jobs = append(jobs, queue.Job{
				Name:   fmt.Sprintf("Importing file to db : %s", f.Name()),
				Action: action,
			})

		}

		//add jobs to queue
		fileToProcess.AddJobs(jobs)

		//define queue worker that will execute our queue
		worker := queue.NewWorker(fileToProcess)

		//execute job in queue
		worker.DoWork()

		//sleep program for 10 sec
		time.Sleep(10 * time.Second)
	}

}
