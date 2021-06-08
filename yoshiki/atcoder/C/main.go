package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type village struct {
	vilnum, money int
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {

	var k, n int

	sc.Split(bufio.ScanWords)

	fmt.Scanf("%d %d", &n, &k)

	var vil []village

	for i := 0; i < n; i++ {

		s := nextInt()

		t := nextInt()

		vil = append(vil, village{vilnum: s, money: t})
		// vil[i].vilnum = a
		// vil[i].money = b
	}

	sort.Slice(vil, func(i, j int) bool { return vil[i].vilnum < vil[j].vilnum })

	//fmt.Println(vil2)

	//k is money

	var current int = 0

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
