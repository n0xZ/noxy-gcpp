package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
)

func printFileModificationsDate() {
	date := time.Now()
	actualDate := fmt.Sprintf("[%d/%d/%d : %d:%d:%d]", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second())
	fmt.Println(actualDate)
}

func main() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	// use goroutine to start the watcher
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// monitor only for write events
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("Modified file:", event.Name)
					executeCompile()
				}
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()
	// tuki
	// provide the file name along with path to be watched
	err = watcher.Add("main.go")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
