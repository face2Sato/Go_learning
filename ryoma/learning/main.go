package main

import f "fmt"

func main() {

	unko := [3]string{"unko", "chinpo", "yoshiki"}

	// var unko [3]string = [3]string{"unko", "chinpo", "yoshiki"}

	// for i := 0; i < len(unko); i++ {
	// 	fmt.Println(unko[i])
	// }

	for _, v := range unko {
		f.Println("map's val:", filterHuman(v))
	}
}

func filterHuman(val string) string {
	if val == "yoshiki" {
		return "Human"
	} else {
		return "Garbage"
	}
}
