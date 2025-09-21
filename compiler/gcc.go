package compiler

import (
	"fmt"
	"os/exec"
)

func compileGCC(fileName string, targetDir string) {
	outputFile := removeFileExtension(fileName)
	printFileModificationDate()
	fmt.Println("⏳ Compilando en GCC:", fileName, "-o", outputFile)
	cmd := exec.Command("g++", fileName, "-o", outputFile)
	rewriteOutputDefaultsToTargetDir(targetDir, cmd)
	err := cmd.Run()
	if err != nil {
		LogError(err)
	}
	printFileModificationDate()
	fmt.Println("\033[32m✅ Compilado con éxito. \033[32m")
}
