package validation

import (
	"regexp"
)

//TODO Validate Rules according to a standard
func ValidateRules(rules string) bool {

	/*
		Excel Properties:
		1. Rows from 1 onwards ([1-9]*)
		2. Columns from A-Z only ([A-Z])
	*/
	regex := ""

	matched, _ := regexp.MatchString(regex, rules)
	return matched

}
