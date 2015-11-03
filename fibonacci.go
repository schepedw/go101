package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var err error
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	fib_sequence := make([]int, n)
	for i := 0; i < n; i++ {
		min_bound := math.Max(float64(i-2), 0.0)
		max_bound := math.Max(float64(i), 0.0)
		fib_sequence[i] = nth_fib(i, fib_sequence[int(min_bound):int(max_bound)])
	}
	fmt.Println(fib_sequence)
}

func nth_fib(fib_index int, fib_sequence []int) (nth_fib int) {
	switch fib_index {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib_sequence[0] + fib_sequence[1]
	}
}
