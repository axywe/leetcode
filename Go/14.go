package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, str := range strs {
		for j := len(prefix); j >= 0; j-- {
			prefix = prefix[:j]
			if strings.HasPrefix(str, prefix) {
				fmt.Println(str, prefix, strings.HasPrefix(str, prefix))
				break
			}
		}
	}
	return prefix
}
