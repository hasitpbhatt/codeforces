package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func fun() {
	var n int64
	scanf("%d\n", &n)
	arr := make([]int64, n)
	for i := range arr {
		if i == int(n-1) {
			scanf("%d\n", &arr[i])
		} else {
			scanf("%d ", &arr[i])
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	mn := arr[0]
	ans := int64(0)
	for i, v := range arr {
		mn = min(mn, v)
		ans = max(ans, min(mn, int64(i+1)))
	}
	printf("%d\n", ans)
}

func main() {
	// Flush stdout
	defer writer.Flush()

	var n int
	scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fun()
	}
}
