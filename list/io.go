package list

import (
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
	return ioutil.WriteFile(path, buf, os.ModePerm)
}

// Load ファイルから読込
func (l *List) Load() error {
	path := getSaveFilePath()
	if isNotExist(path) {
		if err := createFile(path); err != nil {
			return err
		}
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(buf, &l.items)
}

func getSaveFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(homeDir, filename)
}

func isNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func createFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return f.Close()
}
