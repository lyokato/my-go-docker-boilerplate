package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var targetFiles = []string{
	"./build/web/Dockerfile",
	"./build/web/image.sh",
	"./build/web/app.sh",
	"./develop/web/docker-compose.yaml",
}

var targetDirTemplate = "src/%s"

var origin = "myapp"
var modified = "yourapp"

func renameProjectNameInFile(filePath string) error {
	stat, err := os.Stat(filePath)
	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %s", err)
	}
	body := string(input)
	replaced := strings.Replace(body, origin, modified, -1)
	err = ioutil.WriteFile(filePath, []byte(replaced), stat.Mode())
	if err != nil {
		return fmt.Errorf("failed to write file: %s", err)
	}
	return nil
}

func main() {
	curDir, err := os.Getwd()
	root := filepath.Join(curDir, "../../")
	for _, file := range targetFiles {
		err = renameProjectNameInFile(filepath.Join(root, file))
		if err != nil {
			fmt.Println(err)
		}
	}
	err = os.Rename(
		filepath.Join(root, fmt.Sprintf(targetDirTemplate, origin)),
		filepath.Join(root, fmt.Sprintf(targetDirTemplate, modified)))
	if err != nil {
		fmt.Println(err)
	}
}
