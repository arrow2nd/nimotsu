package track

import "fmt"

// View 情報を表示
func (t *Track) View() {
	fmt.Printf("追跡番号 : %s\n", t.number)

	for i, status := range t.statuses {
		fmt.Printf("%d) %s %s %s\n", i+1, status.Message, status.Date, status.Office)
	}
}
