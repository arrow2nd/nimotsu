package cmd

import (
	"fmt"

	"github.com/arrow2nd/nimotsu/pack"
	"github.com/spf13/pflag"
)

// getCarrierName 運送業者名をフラグから取得
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

	if enabledFlagCount != 1 {
		return "", fmt.Errorf("expected exactly one of `--japanpost`, `--yamato`, or `--sagawa` to be true")
	}

	return carrier, nil
}
