package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	sf := strings.Fields(s)
	for i := 0; i < len(sf); i ++{
		_, ok := res[sf[i]]
		if ok {
			res[sf[i]] ++
		}else{
			res[sf[i]] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
