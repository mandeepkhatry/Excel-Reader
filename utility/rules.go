package utility

import (
	"regexp"
	"strings"
)

//Seperate different types of rules and work accordingly. Note : rules here are valid
func SeperateExpressions(rules string) (map[string][]string, error) {
	rulesSet := strings.Split(rules, "&&")

	/*
		Excel Properties:
		1. Rows from 1 onwards ([1-9]*)
		2. Columns from A-Z only ([A-Z])
	*/
	nullRegex := "[!]*IsNull[(]([A-Z]|[1-9]*)[)]"

	rulesMap := make(map[string][]string)

	for _, v := range rulesSet {
		matched, err := regexp.MatchString(nullRegex, v)
		if err != nil {
			return map[string][]string{}, err
		}
		if matched {
			rulesMap["empty"] = append(rulesMap["empty"], v)
		} else {
			rulesMap["operation"] = append(rulesMap["operation"], v)
		}
	}

	return rulesMap, nil
}
