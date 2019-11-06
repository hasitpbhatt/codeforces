package main

// Template
import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func readNumberFromLine(n int64) []int64 {
	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		if i == n-1 {
			scanf("%d\n", &arr[i])
		} else {
			scanf("%d ", &arr[i])
		}
	}
	return arr
}

func possible(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	cnt := 0
	change := make([]int, 0, len(s1))
	for i := range s1 {
		if s1[i] != s2[i] {
			cnt++
			change = append(change,i)
		}
	}
	if cnt == 0 {
		return true
	}
	if cnt != 2{
		return false
	}
	return s1[change[0]] == s1[change[1]] && s2[change[0]] == s2[change[1]]
}

func solve() {
	var n int64
	var s1, s2 string
	scanf("%d\n", &n)
	scanf("%s\n", &s1)
	scanf("%s\n", &s2)
	if possible(s1,s2) {
		printf("YES\n")
	} else {
		printf("NO\n")
	}
}

func main() {
	// Flush stdout
	defer writer.Flush()

	var n int
	scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		solve()
	}
}
