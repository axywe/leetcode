package main

import "strings"

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	arr := []string{}
	for i := 0; i < len(parts); i++ {
		if parts[i] == "." || parts[i] == "" {
			continue
		}
		if parts[i] == ".." {
			if len(arr) == 0 {
				continue
			}
			arr = arr[:len(arr)-1]
			continue
		}
		arr = append(arr, parts[i])
	}
	result := []string{}
	for i := 0; i < len(arr); i++ {
		result = append(result, "/")
		result = append(result, arr[i])
	}
	if len(result) == 0 {
		result = append(result, "/")
	}
	return strings.Join(result, "")
}
