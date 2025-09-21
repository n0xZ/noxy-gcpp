package compiler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func getExecPath() string {
	execPath, _ := os.Getwd()
	return execPath
}
func compileBCC32(fileName string) {
	fileNameWithoutPath := filepath.Base(fileName)
	outputFile := removeFileExtension(fileNameWithoutPath)
	printFileModificationDate()
	fmt.Println("⏳ Compilando en BCC32:", fileName)
	cmd := exec.Command("bcc32", outputFile)

	fmt.Println(getExecPath() + fileNameWithoutPath + ".*")
	getExecPath()

	err := cmd.Run()
	if err != nil {
		LogError(err)
	}
	printFileModificationDate()
	fmt.Println("\033[32m✅ Compilado con éxito. \033[32m")
}
