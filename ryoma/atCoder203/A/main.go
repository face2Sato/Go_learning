package main

import "fmt"

func main() {
	var answer int
	s := []int{0, 0, 0}
	fmt.Scanf("%d %d %d", &s[0], &s[1], &s[2])
	if s[0] == s[1] {
		answer = s[2]
	} else if s[1] == s[2] {
		answer = s[0]
	} else if s[2] == s[0] {
		answer = s[1]
	} else {
		answer = 0
	}

	fmt.Println(answer)
}
