package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Scanf("%d %d %d", &a, &b, &c)

	if a == b {
		fmt.Println(c)
	} else if a == c {
		fmt.Println(b)
	} else if b == c {
		fmt.Println(a)
	} else {
		fmt.Printf("0\n")
	}

}
