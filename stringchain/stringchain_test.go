package stringchain

import "testing"

func stringChainReplacements(a []string) (r int) {
	i := 0
	for i < len(a) {
		n := map[byte]int{a[i][len(a[i])-1]: 1}
		i++
		for i < len(a) {
			x, _ := n[a[i][0]]
			n[a[i][0]] = x + 1
			if len(a[i]) > 1 {
				break
			}
			i++
		}
		m := 0
		for _, x := range n {
			if x > m {
				m = x
			}
			r += x
		}
		r -= m
	}
	return
}

func stringChainReplacements2(a []string) (r int) {
	s := a[0]
	c := s[len(s)-1]
	for i := 1; i < len(a); i++ {
		if len(a[i]) > 1 && a[i][0] == c {
			c = a[i][len(a[i])-1]
			continue
		}
		n := map[byte]int{c: 1}
		for i < len(a) {
			x, _ := n[a[i][0]]
			n[a[i][0]] = x + 1
			if len(a[i]) > 1 {
				break
			}
			i++
		}
		max, total := 0, 0
		for _, x := range n {
			if x > max {
				max = x
			}
			total += x
		}
		r += total - max
		if i < len(a) {
			c = a[i][len(a[i])-1]
		}
	}
	return
}

func test(t *testing.T, a []string, expected int) {
	r := stringChainReplacements(a)
	if r != expected {
		t.Errorf("Expected %d, got %d", expected, r)
	}
}
func Test1(t *testing.T) {
	r := stringChainReplacements([]string{"abc", "def", "ghi"})
	if r != 2 {
		t.Error("Expected 2, got", r)
	}
}

func Test3(t *testing.T) {
	r := stringChainReplacements([]string{"abc", "cde", "efg", "ghi"})
	if r != 0 {
		t.Error("Expected 0, got", r)
	}
}

func Test10(t *testing.T) {
	test(t, []string{"xt", "x", "fqrhsqxwt"}, 2)
	test(t, []string{"tontq", "nk"}, 1)
	test(t, []string{
		"xt",
		"x",
		"fqrhsqxwt",
		"tontq",
		"nk",
		"knplki",
		"i",
		"x",
		"g",
		"gifguc",
		"e",
	}, 7)
}
