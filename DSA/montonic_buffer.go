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
        max  []int
        min  []int
}

func NewMonotonicBuffer(size int) *MonotonicBuffer {
        return &MonotonicBuffer{
                max: make([]int, size),
                min: make([]int, size),
        }
}

func (b *MonotonicBuffer) Add(val int) {
        b.max[0] = max(b.max[0], val)
        b.min[0] = min(b.min[0], val)

        for i := 1; i < len(b.max); i++ {
                b.max[i] = b.max[i-1]
                b.min[i] = b.min[i-1]

                if val > b.max[i] {
                        b.max[i] = val
                }

                if val < b.min[i] {
                        b.min[i] = val
                }
        }
}

func (b *MonotonicBuffer) GetMax() int {
        return b.max[0]
}

func (b *MonotonicBuffer) GetMin() int {
        return b.min[0]
}

func main() {
        buf := NewMonotonicBuffer(3)

        buf.Add(5)
        buf.Add(10)
        buf.Add(15)
        buf.Add(-1)

        fmt.Println("Max:", buf.GetMax())  // Output: Max: 15
        fmt.Println("Min:", buf.GetMin())  // Output: Min: -1
}

