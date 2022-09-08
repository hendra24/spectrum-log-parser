package main

import (
	"log"
	"time"

	"github.com/hendra24/spectrum-log-parser/file_processor"
)

func main() {
	start := time.Now()
	file_processor.ProcessFile()

	elapse := time.Since(start)
	log.Println("duration", elapse.Seconds(), "seconds")
}
