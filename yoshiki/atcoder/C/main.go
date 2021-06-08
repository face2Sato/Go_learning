package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	k       int
	n       int
	current int
	vil     []village
)

type village struct {
	vilnum int
	money  int
}

func main() {

	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	fmt.Scanf("%d %d", &n, &k)
	for i := 0; i < n; i++ {
		s := nextInt(sc)
		t := nextInt(sc)
		vil = append(vil, village{vilnum: s, money: t})
	}

	sort.Slice(vil, func(i, j int) bool { return vil[i].vilnum < vil[j].vilnum })

	for i := 0; i < len(vil); i++ {
		if k-(vil[i].vilnum-current) >= 0 {
			k += vil[i].money - (vil[i].vilnum - current)
			current += (vil[i].vilnum - current)
		} else {
			break
		}
		current = vil[i].vilnum
	}
	current += k

	fmt.Println(current)
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
