package cmd

import (
	"fmt"

	"github.com/arrow2nd/nimotsu/pack"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/pflag"
)

// getCarrierName : 配送業者IDをフラグから取得
func getCarrierName(flags *pflag.FlagSet) (pack.CarrierName, error) {
	for name, key := range pack.GetCarriers() {
		if exist, _ := flags.GetBool(string(key)); exist {
			return name, nil
		}
	}

	// フラグの指定が不正なら選択させる
	return selectCarrier()
}

// inputComment : コメントを入力
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

// selectCarrier : 配送業者を選択
func selectCarrier() (pack.CarrierName, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   `{{ ">" | cyan }} {{ . | cyan }}`,
		Inactive: "  {{ . }}",
		Selected: `{{ "Carrier:" | faint }} {{ . }}`,
	}

	prompt := promptui.Select{
		Label:     "Carrier",
		Items:     pack.GetCarrierNames(),
		Templates: templates,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return pack.CarrierName(result), nil
}

// selectTrackingNumber : リスト内の追跡番号を選択
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
