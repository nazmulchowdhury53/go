package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ { 
		var a, c, d int
		fmt.Scan(&a, &c, &d)
		if a+d == c || a+c == d || d+c == a {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}