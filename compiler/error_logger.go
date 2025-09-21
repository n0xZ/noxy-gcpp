package compiler

import (
	"fmt"
	"log"
)

func LogError(err error) {
	printFileModificationDate()
	fmt.Println("\033[31m🚨 Ocurrió un error: ")
	log.Fatal(err.Error())
	fmt.Print("\033[0m")
}
