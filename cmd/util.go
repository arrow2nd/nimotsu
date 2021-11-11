package cmd

import (
	"fmt"

	"github.com/arrow2nd/nimotsu/pack"
	"github.com/spf13/pflag"
)

// getCarrierName 運送業者名をフラグから取得
func getCarrierName(flags *pflag.FlagSet) (string, error) {
	isJapanPost, _ := flags.GetBool("japanpost")
	isYamato, _ := flags.GetBool("yamato")
	isSagawa, _ := flags.GetBool("sagawa")

	enabledFlagCount := 0
	carrier := ""

	if isJapanPost {
		enabledFlagCount++
		carrier = pack.JapanPost
	}
	if isYamato {
		enabledFlagCount++
		carrier = pack.YamatoTransport
	}
	if isSagawa {
		enabledFlagCount++
		carrier = pack.SagawaExpress
	}

	if enabledFlagCount != 1 {
		return "", fmt.Errorf("expected exactly one of `--japanpost`, `--yamato`, or `--sagawa` to be true")
	}

	return carrier, nil
}
