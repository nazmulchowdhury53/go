package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func compress(s string) string {
	var res strings.Builder
	for i, r := range s {
		lower := unicode.ToLower(r)
		if res.Len() == 0 || rune(res.String()[res.Len()-1]) != lower {
			res.WriteRune(lower)
		}
	}
	return res.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)

		compressed := compress(s)

		if compressed == "meow" {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
