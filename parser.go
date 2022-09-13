package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	queue "github.com/hendra24/spectrum-log-parser/data_structure"
	db_connector "github.com/hendra24/spectrum-log-parser/db"
	"github.com/hendra24/spectrum-log-parser/file_processor"
)

const spectrum_db string = "spectrum_log"

func main() {

	//initialize new queue
	fileToProcess := queue.NewQueue("Parser File to DB")
	//context for process
	ctx := context.Background()
	//connect to db
	db, err := db_connector.Connect(ctx, spectrum_db)
	if err != nil {
		log.Fatal(err)
	}
	//check directory folder have file ? if y do something
	for {

		var jobs []queue.Job
		log.Println("Checking log in path " + string(file_processor.DATA_LOGS_PATH))
		files, err := ioutil.ReadDir(file_processor.DATA_LOGS_PATH)
		if err != nil {
			log.Fatal(err)
		}

		// check if theris file or not in directory
		if len(files) != 0 {

			for _, f := range files {
				if f.IsDir() {
					continue
				}

				file_name := f.Name()

				var action = func() error {

					err = file_processor.ReadFile(ctx, file_name, "\t", db)
					if err != nil {
						return err
					}

					return nil

				}

				jobs = append(jobs, queue.Job{
					Name:   fmt.Sprintf("Job for file %s", f.Name()),
					Action: action,
				})
			}

			//add jobs to queue
			fileToProcess.AddJobs(jobs)

			//define queue worker that will execute our queue
			worker := queue.NewWorker(fileToProcess)

			//execute job in queue
			worker.DoWork()

		} else {
			//DELETE UNSUEFUL LOG
			err = file_processor.DeleteCollection(ctx, db)
			if err != nil {
				log.Println(err.Error())
			} else {
				log.Println("DB SAMPAH DI Hapus")
			}
			// if folder empty print
			log.Println("Directrory empty... no file to process")
			//sleep program for 30 sec
			time.Sleep(30 * time.Second)
		}

	}

}
