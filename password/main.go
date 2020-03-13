package main

import (
	"log"
	"strconv"
)

func main() {
	// 128392-643281
	var count int
	for i := 128392; i <= 643281; i++ {
		s := strconv.Itoa(i)
		if repeatedCharacters(s) && increasing(s) {
			count++
		}
	}

	log.Printf("increasing? %t", increasing("111111"))
	log.Printf("increasing? %t", increasing("223450"))
	log.Printf("repeated count %d\n", count)
}

func increasing(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

func repeatedCharacters(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			// log.Printf("%v == %v", s[i], s[i+1])
			return true
		}
	}
	return false
}
