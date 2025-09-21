package main

import (
	"flag"
	"log"
	"noxy-gcpp/compiler"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	var f, c string
	flag.StringVar(&f, "f", "filename", "Should specify the filename")
	flag.StringVar(&c, "c", "compilerOption", "Should specify the compilerOption (GCC or BCC32)")
	targetDir := filepath.Dir(f)
	flag.Parse()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	var eventTimer *time.Timer
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					if eventTimer != nil {
						eventTimer.Stop()
					}
					eventTimer = time.AfterFunc(100*time.Millisecond, func() {
						compiler.CompileFile(c, f, targetDir)
					})

				}
			case err := <-watcher.Errors:
				compiler.LogError(err)
			}
		}
	}()
	err = watcher.Add(f)
	if err != nil {
		compiler.LogError(err)
	}
	<-done
}
