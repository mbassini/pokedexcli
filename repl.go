package main

import (
	"strings"
)

// My implementation LOL
//func cleanInput(text string) []string {
//	trimmed := strings.TrimSpace(text)
//	var slice []string
//	split := strings.Split(trimmed, " ")
//	for _, v := range split {
//		if len(v) > 0 {
//			slice = append(slice, strings.ToLower(v))
//		}
//	}
//	return slice
//}

// Course impl
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
