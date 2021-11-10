package track

import "fmt"

// View 表示
func (t *Track) View() {
	fmt.Println(t.number)
	fmt.Println(t.statuses)
}
