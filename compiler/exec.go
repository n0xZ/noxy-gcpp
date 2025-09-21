package compiler

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func logError(err error) {
	printFileModificationDate()
	fmt.Println("\033[31m🚨 Ocurrió un error: ")
	log.Fatal(err.Error())
	fmt.Print("\033[0m")
}
func rewriteOutputDefaultsToTargetDir(targetDir string, cmd *exec.Cmd) {
	cmd.Dir = targetDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
}
func CompileFile(compiler string, fileName string, targetDir string) {

	switch compiler {
	case "gcc":
		printFileModificationDate()
		fmt.Println("⏳ Compilando en GCC:", fileName, "-o", "")
		cmd := exec.Command("g++", fileName)
		rewriteOutputDefaultsToTargetDir(targetDir, cmd)
		err := cmd.Run()
		if err != nil {
			logError(err)
		}
		printFileModificationDate()
		fmt.Println("\033[32m✅ Compilado con éxito. \033[32m")

	case "bcc32":
		printFileModificationDate()
		fmt.Println("⏳ Compilando en BCC32:", fileName)
		cmd := exec.Command("bcc32", fileName)
		rewriteOutputDefaultsToTargetDir(targetDir, cmd)
		err := cmd.Run()
		if err != nil {
			logError(err)
		}
		printFileModificationDate()
		fmt.Println("\033[32m✅ Compilado con éxito. \033[32m")

	}

}
