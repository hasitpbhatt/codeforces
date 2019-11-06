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

func solve() {
	// Do something
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
