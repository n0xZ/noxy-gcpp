package compiler

import (
	"fmt"
	"time"
)

func printFileModificationsDate() {
	date := time.Now()
	actualDate := fmt.Sprintf("\033[33;1m [%d/%02d/%02d %02d:%02d:%02d] \033[0m",
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second())
	fmt.Print(actualDate)
}
