package main

import "fmt"

func main() {
	req := []string{"aa", "as", "af", "ag"}
	exist := []string{"as", "ag"}
	var notExist []string

	m := make(map[string]interface{}, len(exist))
	for i, v := range exist {
		m[v] = i
	}
	for _, vv := range req {
		if _, ok := m[vv]; !ok {
			notExist = append(notExist, vv)
		}
	}
	fmt.Println(notExist)
}
