package main

import (
	"log"
	"os"
	"sync"
)

var (
	logs  = make(map[string]*log.Logger)
	files = make(map[string]*os.File)
	mutex sync.Mutex
)

func initLogger(logName string) {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := os.OpenFile("SearchLogs"+logName+".txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file %s: %s", logName, err)
	}

	files[logName] = file
	logs[logName] = log.New(file, logName+": ", log.Ldate|log.Ltime|log.Lshortfile)

}

func logTo(logName, message string) {
	mutex.Lock()
	defer mutex.Unlock()

	if logger, ok := logs[logName]; ok {
		logger.Println(message)
	} else {
		log.Printf("Log %s does not exist\n", logName)
	}
}

func closeLoggers() {
	mutex.Lock()
	defer mutex.Unlock()

	for _, file := range files {
		file.Close()
	}
}
