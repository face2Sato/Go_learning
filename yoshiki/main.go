package main

import "fmt"

func main() {

	/*var unko [3]string = [3]string{"unko", "pee", "poop"}

	unko := [3]string{"syamu", "pee", "poop"}

	for _, v := range unko {
		fmt.Println("value = " + v)
		fmt.Println(FindGarbage(v))
	}*/

	num := [11]int{0, 14, 22, 2, 20, 5, 211, 7, 8, 90, 50}
	/*
		for _, v := range num {
			a, b := oddeven(&v)
			fmt.Println(b)

			if a {
				fmt.Printf("hoge\n")
			} else {
				fmt.Printf("hage\n")
			}
		}
	*/
	reverse(num, len(num))
}

func reverse(a [11]int, length int) {

	for i := 0; i < length; i++ {
		defer fmt.Println(a[i])
	}

}

func oddeven(a *int) (bool, string) {
	if *a%2 != 0 {
		return false, "even"
	} else {
		return true, "odd"
	}
}

func FindGarbage(v string) string {
	if v == "syamu" {
		return "garbage"
	} else {
		return "go to toilet"
	}
}

//modified in main
