// export GOEXPERIMENT=rangefunc
package main

import "fmt"

func main() {
	for v := range LinearSpace(0, 3, 10) {
		fmt.Printf("%.3f ", v)
	}
	fmt.Println()
	// 0.000 0.333 0.667 1.000 1.333 1.667 2.000 2.333 2.667 3.000
}

// LinearSpace returns a iterator that yields `count` equally distanced values between `start` and `end` (inclusive).
func LinearSpace(start, end float64, count int) func(func(float64) bool) {
	// Can't return error from iterator.
	if count <= 0 {
		panic(fmt.Sprintf("%d - invalid count", count))
	}

	if start >= end {
		panic(fmt.Sprintf("start (%f) >= end (%f)", start, end))
	}

	// Run count-1 iterations and return `end` last to avoid floating point rounding errors.
	step := (end - start) / float64(count-1)

	fn := func(yield func(float64) bool) {
		n := start
		for i := 0; i < count-1; i++ {
			if !yield(n) {
				return
			}
			n += step
		}

		yield(end)
	}

	return fn
}
