package compiler

import (
	"os"
	"os/exec"
	"strings"
)

func rewriteOutputDefaultsToTargetDir(targetDir string, cmd *exec.Cmd) {
	cmd.Dir = targetDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
}

func removeFileExtension(filename string) string {
	return strings.TrimSuffix(filename, ".cpp")
}
func CompileFile(compiler string, fileName string, targetDir string) {
	switch compiler {
	case "gcc":
		compileGCC(fileName, targetDir)

	case "bcc32":
		compileBCC32(fileName)

	}

}
