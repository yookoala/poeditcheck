package csanitize

import (
	"regexp"
	"strings"
)

var placeholderRe *regexp.Regexp

func init() {
	placeholderRe = regexp.MustCompile("%(|\\d+\\$)[bcdeEfFgGosuxX]")
}

// GetTrims get strings from both end that consists only
// with character in cutset.
func GetTrims(s, cutset string) (left, right string) {
	l := len(s)
	leftLength := l - len(strings.TrimLeft(s, cutset))
	rightBegin := len(strings.TrimRight(s, cutset))
	left, right = s[:leftLength], s[rightBegin:]
	return
}

// GetPlaceholder gets the C-style placeholders
// found in the given string.
func GetPlaceholder(s string) []string {
	return placeholderRe.FindAllString(s, -1)
}
