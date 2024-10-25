// 242. Valid Anagram https://leetcode.com/problems/valid-anagram/description/
package main

import "fmt"

func main() {
	input := "nagaram"

	input1 := "nagaram"
	fmt.Println(isAnagram(input, input1))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false // Early exit if lengths differ
	}
	smap := make(map[rune]int)

	tmap := make(map[rune]int)

	for _, char := range s {
		smap[char]++
	}
	for _, char := range t {
		tmap[char]++
	}

	if compare(smap, tmap) {
		return true
	}
	return false
}

// function to compare if the two maps have similar key-value pairs
func compare(m1, m2 map[rune]int) bool {
	for key, value := range m1 {
		if v, exists := m2[key]; !exists || v != value {
			return false
		}
	}

	for key, value := range m2 {
		if v, exists := m1[key]; !exists || v != value {
			return false
		}
	}

	return true
}
