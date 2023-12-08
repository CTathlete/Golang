package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func CopyFile(src string, dst string) {
	// Reads src and writes it to dst.
	data, err := os.ReadFile(src)
	CheckErr(err)
	err = os.WriteFile(dst, data, 0644)
	CheckErr(err)
}

/*
1. Put nil as conda env name if you're not using an env.
2. PyArgs are the arguments for your python script.
*/
func RunPythonFile(condaEnvNamePtr *string, pyFileDirPtr *string, pyFileNamePtr *string, pyArgs ...string) {
	powerShellCommand := ""

	switch condaEnvNamePtr {
	case nil:
		powerShellCommand = "python" +
			*pyFileDirPtr + "/" + *pyFileNamePtr + strings.Join(pyArgs, " ")
	default:
		powerShellCommand = "conda activate " + *condaEnvNamePtr + ";python" +
			*pyFileDirPtr + "/" + *pyFileNamePtr + strings.Join(pyArgs, " ")
	}

	cmd := exec.Command("powershell", powerShellCommand)
	out, err := cmd.Output()
	fmt.Println(cmd.String()) // Printing the command itself in string form, mainly for debugging purposes.
	fmt.Println(string(out))  // Printing stdout.
	CheckErr(err)
}
