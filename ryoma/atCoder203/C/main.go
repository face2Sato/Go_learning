package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	N          int
	K          int
	_tempArray [][]int
	answer     int
	current    int
)

func main() {
	// - fetch data
	var sc = bufio.NewScanner(os.Stdin)

	// N,K
	if sc.Scan() {
		_tmp := strings.Fields(sc.Text())
		N = unwrapAtoi(_tmp[0])
		K = unwrapAtoi(_tmp[1])
	}

	for i := 0; i < N; i++ {
		_tempArray = append(_tempArray, []int{})
		if sc.Scan() {
			_tmp := strings.Fields(sc.Text())
			_tempArray[i] = append(_tempArray[i], unwrapAtoi(_tmp[0]))
			_tempArray[i] = append(_tempArray[i], unwrapAtoi(_tmp[1]))
		}
	}
	// multiplications
	answer = run()
	fmt.Println(answer)
}

func run() int {
	// sorting stuff
	sort.Slice(_tempArray, func(i, j int) bool { return _tempArray[i][0] < _tempArray[j][0] })

	// every array multiplications
	for j, s := range _tempArray {
		if K < s[0]-current {
			return current + K
		} else {
			// go to the village to get the monimoni
			K -= s[0] - current
			current += s[0] - current
			// get monimoni
			K += _tempArray[j][1]
		}
	}
	return current + K
}

func unwrapAtoi(word string) int {
	_tmp, _ := strconv.Atoi(word)
	return _tmp
}
