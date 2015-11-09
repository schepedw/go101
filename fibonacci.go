package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"path"
	"strconv"
)

func main() {
	fib_server()
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

func fib_server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url_path, fib_max := path.Split(r.URL.Path[1:])
		n, err := strconv.Atoi(fib_max)
		if url_path != "fibonacci/" {
			panic("format url as fibonacci/:n")
		}
		if err != nil {
			panic(err)
		}
		fib_sequence := make([]int, n)
		for i := 0; i < n; i++ {
			min_bound := math.Max(float64(i-2), 0.0)
			max_bound := math.Max(float64(i), 0.0)
			fib_sequence[i] = nth_fib(i, fib_sequence[int(min_bound):int(max_bound)])
		}
		fmt.Fprintln(w, fib_sequence)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
