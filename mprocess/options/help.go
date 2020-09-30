package options

import "flag"

// HELP 帮助变量
var HELP bool

// HelpOptions 帮助选项
type HelpOptions struct {
}

// Help 1234
func (h *HelpOptions) Help() {
	flag.Usage()
}
