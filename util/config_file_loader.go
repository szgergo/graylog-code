package util

import (
	"bufio"
	"com/github/graylog-code/error"
	"os"
	"path/filepath"
	"strings"
)

func LoadConfigurationFile(filePath string) []byte {
	file, err := os.Open(filePath)
	absolutePath := filepath.Dir(filePath)
	error.Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line,"$_ref:") {
			text = append(text, string(LoadConfigurationFile(absolutePath + "/" + strings.Split(line,":")[1])))
		} else {
			text = append(text, line + "\n")
		}
	}

	defer CloseResource(file)

	return []byte(strings.Join(text,""))
}