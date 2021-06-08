package main

import "fmt"

func main() {
	var n, k, sum int

	fmt.Scanf("%d %d", &n, &k)
	// fmt.Printf("unko?")
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			sum += 100*i + j
		}
	}

	fmt.Println(sum)

}
