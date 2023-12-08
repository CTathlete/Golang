package main

import (
	"flag"
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
1. Put "" as conda env name if you're not using an env.
2. PyArgs are the arguments for your python script.
*/
func RunPythonFile(condaEnvNamePtr *string,
	pyFileDirPtr *string, pyFileNamePtr *string, pyArgs ...string) {
	powerShellCommand := ""

	switch *condaEnvNamePtr {
	case "":
		powerShellCommand = "python3 " +
			*pyFileDirPtr + "/" + *pyFileNamePtr + " " + strings.Join(pyArgs, " ")
	default:
		powerShellCommand = "conda activate " + *condaEnvNamePtr + ";python3 " +
			*pyFileDirPtr + "/" + *pyFileNamePtr + " " + strings.Join(pyArgs, " ")
	}

	cmd := exec.Command("powershell", powerShellCommand)
	out, err := cmd.Output()
	fmt.Println(cmd.String()) // Printing the command itself in string form, mainly for debugging purposes.
	fmt.Println(string(out))  // Printing stdout.
	CheckErr(err)
}

type argList []string

func (a *argList) Set(value string) error {
	*a = append(*a, value)
	return nil
}

func (a *argList) String() string {
	return strings.Join(*a, " ")
}

func main() {
	condaEnvNamePtr := flag.String("cen", "", "Name of conda env, if you're using one.")
	pyFileDirPtr := flag.String("pyd", "", "Path to the python script you'd likr to run.")
	pyFileNamePtr := flag.String("pyn", "",
		"Name of the python script, including suffix.")
	pyArgs := argList{""}
	flag.Var(&pyArgs, "pargs", "Arguments for the python script.")
	flag.Parse()
	RunPythonFile(condaEnvNamePtr, pyFileDirPtr, pyFileNamePtr, pyArgs...)
}
