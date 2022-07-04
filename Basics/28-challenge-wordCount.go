package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)

	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		_, ok := ret[words[i]]
		if ok {
			ret[words[i]] += 1
		} else {
			ret[words[i]] = 1
		}
	}

	return ret
}

func main() {
	wc.Test(WordCount)
}
