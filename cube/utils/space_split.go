package utils

import "strings"

// SpaceSplit splits a string by the SPACE character, but does take into account multiple spaces.
func SpaceSplit(Data string) []string {
	Split := strings.Split(Data, " ")
	SplitNew := make([]string, 0)
	for _, v := range Split {
		if v != "" && v != " " {
			SplitNew = append(SplitNew, v)
		}
	}
	return SplitNew
}
