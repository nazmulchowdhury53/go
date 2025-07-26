package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)

	maxTeamSize := 0
	j := 0 

	for i := 0; i < n; i++ {
		for a[i]-a[j] > 5 {
			j++  
		}
		if i-j+1 > maxTeamSize {
			maxTeamSize = i - j + 1
		}
	}

	fmt.Println(maxTeamSize)
}
