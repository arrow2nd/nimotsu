package list

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	ls := New()

	d := t.TempDir()
	ls.SetDir(d)

	ls.AddItem(&Item{
		Carrier: "Test Express",
		Number:  "0123456",
		Comment: "test",
	})

	err := ls.Save()
	assert.NoError(t, err)

	bytes, err := os.ReadFile(filepath.Join(d, fileName))
	assert.NoError(t, err)

	want := `- carrier: Test Express
  number: "0123456"
  comment: test
`

	assert.Equal(t, want, string(bytes))
}

func TestLoad(t *testing.T) {
	t.Run("読み込めるか", func(t *testing.T) {
		yml := `- carrier: Test Express
  number: "0123456"
  comment: test
`

		d := t.TempDir()
		os.WriteFile(filepath.Join(d, fileName), []byte(yml), os.ModePerm)

		ls := New()
		ls.SetDir(d)

		err := ls.Load()
		assert.NoError(t, err)

		assert.Len(t, ls.items, 1)
		assert.Equal(t, "Test Express", ls.items[0].Carrier)
		assert.Equal(t, "0123456", ls.items[0].Number)
		assert.Equal(t, "test", ls.items[0].Comment)
	})

	t.Run("ファイルがない場合に作成されるか", func(t *testing.T) {
		ls := New()

		d := t.TempDir()
		ls.SetDir(d)

		err := ls.Load()
		assert.NoError(t, err)

		files, err := os.ReadDir(d)
		assert.NoError(t, err)

		assert.Len(t, files, 1)

		for i, f := range files {
			if name := f.Name(); name != fileName {
				t.Errorf("files[%d] : %s != %s", i, name, fileName)
			}
		}
	})
}
