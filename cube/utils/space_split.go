package utils

import "strings"

// SpaceSplit splits a string by the SPACE character, but does take into account multiple spaces.
func SpaceSplit(Data string) []string {
	Split := strings.Split(Data, " ")
	SplitNew := make([]string, len(Split))
	BadStrings := 0
	for i, v := range Split {
		if v == "" {
			BadStrings++
		} else {
			SplitNew[i-BadStrings] = v
		}
	}
	return SplitNew[BadStrings:]
}
