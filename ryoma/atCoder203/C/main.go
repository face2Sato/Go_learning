package main

import (
	"fmt"
	"math"
)

var (
	N       int
	K       int
	A       []int
	B       []int
	C       []bool
	answer  int
	current int
)

func main() {
	// データを取ってくる
	fmt.Scanf("%d %d", &N, &K)
	for i := 0; i < N; i++ {
		_a := 0
		_b := 0
		fmt.Scanf("%d %d", &_a, &_b)
		A = append(A, _a)
		B = append(B, _b)
		C = append(C, false)
	}
	//　実行
	answer = run()

	fmt.Println(answer)
}

func run() int {
	// お金をもらう予定の人数を入れる
	road := len(A)

	for {
		min := math.MaxInt32
		for i, s := range A {
			if C[i] == false && min > s {
				min = s
			}
		}

		for j, s := range A {
			if s == min {
				if K < s-current {
					return current + K
				} else {
					// go to the village to get the monimoni
					K -= s - current
					current += s - current
					// get monimoni
					K += B[j]
				}
				C[j] = true
				road--
			}
		}
		if road == 0 {
			return current + K
		}
	}
}
