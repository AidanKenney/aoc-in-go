package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
	"strconv"
	"unicode"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {

	conversionMap := map[string]string{
		"one": "1", "two": "2", "three": "3", 
		"four": "4", "five": "5", "six": "6", 
		"seven": "7", "eight": "8", "nine": "9"}

	splitLines := strings.Split(input,"\n")

	var sum = 0
	for _, list := range splitLines {
		for k, v := range conversionMap {
			if (strings.Contains(list, k)) {
				// include the first/last parts of number when replacing string
				// there could be other numbers in the string that use those letters
				var sb strings.Builder
				sb.WriteRune(rune(k[0]))
				sb.WriteString(v)
				sb.WriteRune(rune(k[len(k)-1]))
				list = strings.Replace(list, k, sb.String(), -1)
			}
		}
		var sb strings.Builder
		// find first digit
		// cast to rune because backwards iteration is weird in Go
		rs := []rune(list)
		for _, ch := range rs {
			if (unicode.IsDigit(ch)) {
				sb.WriteRune(ch)
				break
			}
		}
		// reverse order
		max := len(rs) - 1
		for i := range rs {
			ch := rs[max-i]
			if (unicode.IsDigit(ch)) {
				sb.WriteRune(ch)
				break
			}
		}
		s1, err1 := strconv.Atoi(sb.String());
		if err1 == nil {
			fmt.Println(s1)
			sum += s1
		}
	}
	return sum
}
