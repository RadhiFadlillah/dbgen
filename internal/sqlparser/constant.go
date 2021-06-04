package sqlparser

import "regexp"

var (
	rxQueryProps = regexp.MustCompile(`(?i)([^\s:,]+)\s*:\s*([^\s:,]+)`)
)
