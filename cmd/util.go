package cmd

import (
	"fmt"

	"github.com/arrow2nd/nimotsu/pack"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/pflag"
)

// getCarrierName 配送業者名をフラグから取得
func getCarrierName(flags *pflag.FlagSet) (string, error) {
	enabledFlagCount := 0
	carrier := ""

	if jp, _ := flags.GetBool("japanpost"); jp {
		enabledFlagCount++
		carrier = pack.JapanPost
	}
	if ym, _ := flags.GetBool("yamato"); ym {
		enabledFlagCount++
		carrier = pack.YamatoTransport
	}
	if sg, _ := flags.GetBool("sagawa"); sg {
		enabledFlagCount++
		carrier = pack.SagawaExpress
	}

	// フラグの指定が不正なら選択させる
	if enabledFlagCount != 1 {
		return selectCarrier()
	}

	return carrier, nil
}

// inputComment コメントを入力
func inputComment() (string, error) {
	prompt := promptui.Prompt{
		Label: "Comment",
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	// コメントがない場合指定の文字列に置き換える
	if result == "" {
		result = noCommentMessage
	}

	return result, nil
}

// selectCarrier 配送業者を選択
func selectCarrier() (string, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   `{{ ">" | cyan }} {{ . | cyan }}`,
		Inactive: "  {{ . }}",
		Selected: `{{ "Carrier:" | faint }} {{ . }}`,
	}

	prompt := promptui.Select{
		Label:     "Carrier",
		Items:     []string{pack.JapanPost, pack.YamatoTransport, pack.SagawaExpress},
		Templates: templates,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

// selectTrackingNumber リスト内の追跡番号を選択
func (c *Cmd) selectTrackingNumber() (string, error) {
	items := c.list.Get()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   `{{ ">" | cyan }} {{ .Number | cyan }}`,
		Inactive: "  {{ .Number }}",
		Selected: `{{ "Tracking number:" | faint }} {{ .Number }}`,
		Details: `{{ "Carrier:" | faint }} {{ .Carrier }}
{{ "Comment:" | faint }} {{ .Comment }}`,
	}

	prompt := promptui.Select{
		Label:     "Tracking number",
		Items:     items,
		Templates: templates,
	}

	idx, _, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return items[idx].Number, nil
}

// showSuccessMessage 完了メッセージ
func showSuccessMessage(text string) {
	fmt.Printf("%s %s\n", color.GreenString("✔"), text)
}
