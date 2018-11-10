package box3

import (
	"fmt"
	"sort"
	"testing"
)

func optimalRotatingBox(w []string, r string) int {
	n, m := len(w), len(w[0])
	n1 := 2*n + 2*m
	a := []int{}
	for x := 0; x < m; x++ {
		a = append(a, (n-1)*m+x)
	}
	for x := n - 1; x >= 0; x-- {
		a = append(a, x*m+m-1)
	}
	for x := m - 1; x >= 0; x-- {
		a = append(a, x)
	}
	for x := 0; x < n; x++ {
		a = append(a, x*m)
	}
	c := make([]int, n*m)
	for i, x := range r {
		c[a[i%n1]] += int('0' - x)
	}
	sort.Ints(c)
	b := []int{}
	for _, s := range w {
		for _, c := range s {
			b = append(b, int(c-'0'))
		}
	}
	sort.Ints(b)
	d := 0
	for i, x := range c {
		d -= x * b[i]
	}
	return d
}

func TestOpt5(t *testing.T) {
	result := optimalRotatingBox([]string{"4", "0", "5", "8", "2"}, "58552851124351775074")
	if result != 242 {
		t.Fatal(result)
	}
}

func TestOpt1(t *testing.T) {
	result := optimalRotatingBox([]string{"01", "21", "10"}, "39513695380152438476")
	if result != 49 {
		t.Fatal(result)
	}
}
