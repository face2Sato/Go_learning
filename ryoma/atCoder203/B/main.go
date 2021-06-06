package main

import "fmt"

func main() {
	var answer int
	_s := []int{0, 0}
	fmt.Scanf("%d %d", &_s[0], &_s[1])

	for j := 0; j < _s[0]; j++ {
		for i := 0; i < _s[1]; i++ {
			answer += (j+1)*100 + (i + 1)
		}
	}

	fmt.Println(answer)
}
