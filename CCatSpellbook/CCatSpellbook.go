package CCatSpellbook

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
	// Reads src (filepath) and writes it to dst (filepath).
	data, err := os.ReadFile(src)
	CheckErr(err)
	err = os.WriteFile(dst, data, 0644)
	CheckErr(err)
}

/*
1. Put "" as conda env name if you're not using an env.
2. PyArgs are the arguments for your python script.
*/
func RunPythonFile(condaEnvNamePtr *string, pySuffix string,
	pyFileDirPtr *string, pyFileNamePtr *string, pyArgs ...string) {
	powerShellCommand := ""

	switch *condaEnvNamePtr {
	case "":
		powerShellCommand = "python" + pySuffix + " " +
			*pyFileDirPtr + "/" + *pyFileNamePtr + " " + strings.Join(pyArgs, " ")
	default:
		powerShellCommand = "conda activate " + *condaEnvNamePtr + ";python" + pySuffix + " " +
			*pyFileDirPtr + "/" + *pyFileNamePtr + " " + strings.Join(pyArgs, " ")
	}

	cmd := exec.Command("powershell", powerShellCommand)
	out, err := cmd.Output()
	fmt.Println(cmd.String()) // Printing the command itself in string form, mainly for debugging purposes.
	fmt.Println(string(out))  // Printing stdout.
	CheckErr(err)
}

func FileToString(filePath string) string {
	data, err := os.ReadFile(filePath)
	CheckErr(err)
	return string(data)
}

func StringToFile(data string, targetPath string) int {
	// Opens file for reading AND writing.
	file, err := os.OpenFile(targetPath, os.O_RDWR, 0777)
	CheckErr(err)
	defer file.Close()

	bytesWritten, err := file.WriteString(data)
	CheckErr(err)
	return bytesWritten
}

func ReplaceInFile(oldString string, newString string,
	filePath string, numOfOccurences int) {
	data := FileToString(filePath)
	data = strings.Replace(data, oldString, newString, numOfOccurences)
	StringToFile(data, filePath)
}

type argList []string

func (a *argList) Set(value string) error {
	*a = append(*a, value)
	return nil
}

func (a *argList) String() string {
	return strings.Join(*a, " ")
}

// func main() {
// 	condaEnvNamePtr := flag.String("cen", "", "Name of conda env, if you're using one.")
// 	pyFileDirPtr := flag.String("pyd", "", "Path to the python script you'd likr to run.")
// 	pyFileNamePtr := flag.String("pyn", "",
// 		"Name of the python script, including suffix.")
// 	pyArgs := argList{""}
// 	flag.Var(&pyArgs, "pargs", "Arguments for the python script.")
// 	flag.Parse()
// 	RunPythonFile(condaEnvNamePtr, "3", pyFileDirPtr, pyFileNamePtr, pyArgs...)
// }
