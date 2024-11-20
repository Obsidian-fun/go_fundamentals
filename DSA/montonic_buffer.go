/***
A monotonically increasing buffer (also known as a monotonic queue or stack) is a data structure that
maintains the maximum (or minimum) value seen so far. It's often used to track the maximum (or minimum)
value encountered in an iteration, allowing you to compute statistics like max/min values over time.

***/

package main

import (
        "fmt"
)

type MonotonicBuffer struct {
        _max  []int
        _min  []int
}

func NewMonotonicBuffer(size int) *MonotonicBuffer {
        return &MonotonicBuffer{
                _max: make([]int, size),
                _min: make([]int, size),
        }
}

func (b *MonotonicBuffer) Add(val int) {
        b._max[0] = max(b._max[0], val)
        b._min[0] = min(b._min[0], val)
}

func (b *MonotonicBuffer) GetMax() int {
        return b._max[0]
}

func (b *MonotonicBuffer) GetMin() int {
        return b._min[0]
}

func main() {
        buf := NewMonotonicBuffer(3)

        buf.Add(5)
        buf.Add(10)
        buf.Add(15)
        buf.Add(-1)
				buf.Add(45)
				buf.Add(-23)

        fmt.Println("Max:", buf.GetMax())  // Output: Max: 15
        fmt.Println("Min:", buf.GetMin())  // Output: Min: -1
}

