package mystring

import (
	"fmt"
	"sort"
)

func StringAlphapitic([]string) []string {
	var sc []string
	sort.Strings(sc)
	fmt.Printf("Array in alphabetic order:", "%v", sc)
	return sc
}
