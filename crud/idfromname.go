package crud

import (
	"regexp"
	"strings"
)


var allButLettersRegex regexp.Regexp

func init() {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		panic(err)
	}
	allButLettersRegex = *reg
}

func filterString(str string) string {
    return allButLettersRegex.ReplaceAllString(str, "")
}

func CreateIDFromName(name string) string {
	id := strings.ToLower(name)
	id = filterString(id)
	return id
}
