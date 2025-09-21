package compiler

import (
	"fmt"
	"time"
)

func printFileModificationsDate() {
	date := time.Now()
	actualDate := fmt.Sprintf("[%d/%02d/%02d %02d:%02d:%02d]",
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second())
	fmt.Print(actualDate)
}
