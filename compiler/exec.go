package compiler

import (
	"fmt"
	"log"
	"os/exec"
)

func CompileFile(compiler string, fileName string) {
	fmt.Println("Compilando ", fileName, "...")

	switch compiler {
	case "gcc":

		cmd := exec.Command("g++", fileName)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Compilado con éxito.")

	case "bcc32":
		formattedCommand := fmt.Sprintf("bcc32 %s", fileName)
		cmd := exec.Command(formattedCommand)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		printFileModificationsDate()
		print("Compilado con éxito.")

	}

}
