package regexputil

import (
	"regexp"
)

// Match return first submatch
// return "" if match error or not match
// Example:
//     Match(`Foo(.+$)`, "Foobar") // return "bar", true
func Match(pattern string, s string) (ret string, matched bool) {
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return
	}
	mcs := rx.FindStringSubmatch(s)
	if len(mcs) >= 2 {
		return mcs[1], true
	}
	return "", false
}

// IsMatch return true if match for pattern
func IsMatch(pattern string, s string) bool {
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false
	}
	return matched
}
