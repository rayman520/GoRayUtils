package rayUtils

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func CreateDir(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			rayUtilsGlobals.Logger.Log(LOGERROR, "Could not create directory. %v", err)
			return false
		}
		rayUtilsGlobals.Logger.Log(LOGINFO, "Directory %s created", path)
	}
	return true
}

func MoveFile(inpath, outpath string) {
	CreateDir(outpath)
	cmd := exec.Command("mv", inpath, outpath)
	ExecPrintCmd(cmd)
}

func ReadFileContent(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), err
}

func WriteFileContent(path string, content string) bool {
	if CreateDir(filepath.Dir(path)) {
		file, err := os.Create(path)
		if err != nil {
			rayUtilsGlobals.Logger.Log(LOGERROR, "Could not create file %s. %v", path, err)
			return false
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		writer.WriteString(content)
		writer.Flush()
		rayUtilsGlobals.Logger.Log(LOGINFO, "%s written.", path)
		return true
	}
	return false
}
