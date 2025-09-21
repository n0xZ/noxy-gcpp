package main

import (
	"flag"
	"log"
	"noxy-gcpp/compiler"

	"github.com/fsnotify/fsnotify"
)

func main() {
	var f, c string
	flag.StringVar(&f, "f", "filename", "Should specify the filename")
	flag.StringVar(&c, "c", "compilerOption", "Should specify the compilerOption (GCC or BCC32)")
	flag.Parse()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					compiler.CompileFile(c, f)

				}
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()
	err = watcher.Add(f)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
