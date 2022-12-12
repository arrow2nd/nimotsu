package list

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const fileName = ".nimotsu"

// SetDir : ディレクトリを設定
func (l *List) SetDir(dir string) {
	l.dir = dir
}

// Save : ファイルへ保存
func (l *List) Save() error {
	buf, err := yaml.Marshal(l.packages)
	if err != nil {
		return err
	}

	path := filepath.Join(l.dir, fileName)
	return os.WriteFile(path, buf, os.ModePerm)
}

// Load : ファイルから読込
func (l *List) Load() error {
	path := filepath.Join(l.dir, fileName)
	if isNotExist(path) {
		if err := createFile(path); err != nil {
			return err
		}
	}

	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(buf, &l.packages)
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
