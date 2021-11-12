package list

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Save ファイルへ保存
func (l *List) Save() error {
	buf, err := yaml.Marshal(l.items)
	if err != nil {
		return err
	}

	path := getSaveFilePath()
	err = ioutil.WriteFile(path, buf, os.ModePerm)
	if err != nil {
		fmt.Println("Error: Failed to write file")
		panic(err)
	}

	return nil
}

// Load ファイルから読込
func (l *List) Load() error {
	path := getSaveFilePath()
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &l.items)
	if err != nil {
		return err
	}

	return nil
}

func getSaveFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(homeDir, filename)
}
