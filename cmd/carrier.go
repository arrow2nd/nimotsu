package cmd

type carrier struct {
	isJapanPost bool
	isYamato    bool
	isSagawa    bool
}

func (c *carrier) isValid() bool {
	enabledFlagCount := 0

	if c.isJapanPost {
		enabledFlagCount++
	}
	if c.isYamato {
		enabledFlagCount++
	}
	if c.isSagawa {
		enabledFlagCount++
	}

	return enabledFlagCount == 1
}
