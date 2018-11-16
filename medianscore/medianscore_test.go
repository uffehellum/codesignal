package medianscore

import (
	"container/heap"
	"testing"
)

type ih []int

func (h ih) Len() int            { return len(h) }
func (h ih) Less(i, j int) bool  { return h[i] < h[j] }
func (h ih) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ih) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *ih) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func medianScores(scores []int) (r []int) {
	var low, high ih
	heap.Init(&low)
	heap.Init(&high)
	for _, x := range scores {
		heap.Push(&low, -x)
		if len(low)-len(high) > 0 {
			heap.Push(&high, -heap.Pop(&low).(int))
		}
		if len(low)-len(high) < 0 {
			heap.Push(&low, -heap.Pop(&high).(int))
		}
		if len(low) > len(high) {
			r = append(r, -low[0])
		} else {
			r = append(r, (high[0]-low[0]+1)/2)
		}
	}
	return
}

func test(t *testing.T, scores, expected []int) {
	result := medianScores(scores)
	if len(expected) != len(result) {
		t.Error("wrong count. Got", result, "expected", expected)
		return
	}
	for i := range expected {
		if expected[i] != result[i] {
			t.Error("wrong value pos", i, "got", result, "expected", expected)
		}
	}
}

func Test1(t *testing.T) {
	test(t, []int{100, 20, 50, 70, 45}, []int{100, 60, 50, 60, 50})
}

func Test2(t *testing.T) {
	test(t, []int{10, 20, 30}, []int{10, 15, 20})
}

func Test3(t *testing.T) {
	test(t, []int{98, 91, 70, 26, 75, 91, 30, 88, 86}, []int{98, 95, 91, 81, 75, 83, 75, 82, 86})
}
