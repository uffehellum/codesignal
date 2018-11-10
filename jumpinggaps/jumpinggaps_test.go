package jumpinggaps

import (
	"fmt"
	"strconv"
	"testing"
)

func jumpingGaps(stage []string) int {
	h, w := len(stage), len(stage[0])
	world := make([]bool, w)
	stepsTo := make(map[int]int)
	start, end, p := 0, 0, w
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			ch := stage[r][c]
			world = append(world, ch == '#')
			if ch == 'S' {
				start = p
			}
			if ch == 'E' {
				end = p
			}
			p++
		}
	}
	p, q := 0, []int{start}
	stepsTo[start] = 0
	for len(q) > 0 {
		p, q = q[0], q[1:]
		next := stepsTo[p] + 1
		for _, t := range available(p, world, w) {
			old, ok := stepsTo[t]
			if ok && old <= next {
				continue
			}
			stepsTo[t] = next
			q = append(q, t)
		}
	}
	p, ok := stepsTo[end]
	if !ok {
		p = -1
	}
	print(stage, stepsTo)
	return p
}

func available(p int, world []bool, w int) []int {
	r := jumps(p, world, w)
	s := map[int]bool{}
	for _, j := range r {
		t := drop(j, world, w)
		if t >= 0 {
			s[t] = true
		}
	}
	r = []int{}
	for k, _ := range s {
		r = append(r, k)
	}
	return r
}

func jumps(p int, world []bool, w int) (r []int) {
	c := p % w
	for dr := 0; dr < 4; dr++ {
		t := p - w*dr
		if t > 0 && world[t] {
			break
		}
		for dc := 1; dc < 4; dc++ {
			if c+dc >= w || t > 0 && world[t+dc] {
				break
			}
			r = append(r, t+dc)
		}
		for dc := -1; dc > -4; dc-- {
			if c+dc < 0 || t > 0 && world[t+dc] {
				break
			}
			r = append(r, t+dc)
		}
	}
	return
}

func drop(p int, world []bool, w int) int {
	for p < 0 {
		p += w
	}
	if world[p] {
		return -1
	}
	for p < len(world) && !world[p] {
		p += w
	}
	// if p >= w*h {
	// 	p = -1
	// }
	return p - w
}

func print(world []string, stepsTo map[int]int) {
	w := len(world[0])
	m := [][]rune{make([]rune, w)}
	for _, s := range world {
		m = append(m, []rune(s))
	}
	for k, v := range stepsTo {
		s := strconv.Itoa(v)
		r, c := k/w, k%w
		for i := len(s) - 1; i >= 0; i-- {
			m[r][c] = rune(s[i])
			r--
			if r < 0 {
				break
			}
		}
	}
	for _, a := range m {
		fmt.Println(string(a))
	}
}

func Test8(t *testing.T) {
	result := jumpingGaps([]string{
		"      #   ",
		"          ",
		"   #    #E",
		"   #    ##",
		"    #     ",
		"S   #     ",
		"##  # #   ",
		"    #     ",
		"    #    #",
		"  ###     ",
		"          ",
		"      #   ",
		" # #      ",
		"          ",
		"          ",
		"          ",
	})
	if result != 10 {
		t.Fatal(result)
	}
}
func Test7(t *testing.T) {
	result := jumpingGaps([]string{
		"                                                            ",
		"                                             #####       S  ",
		"                       ##########    ########    #########  ",
		"          ####        ##         # ##                    #  ",
		"          #  ##     ##            #                      #  ",
		"         #     ######       ####  #                      #  ",
		"      ###                  #    # #      ############    #  ",
		"   ####                  ##     #   #####                #  ",
		"   #        ##         ##       #   #                    #  ",
		"          ## #####   ###        #  #                     #  ",
		"          #      #####           ##       #####        ###  ",
		"         ## E           ##  #          ####         ####    ",
		"##     ##   ####        #    #                  #####       ",
		" #### ##          ####  #     #         #  #  ###           ",
		"     ##                 #      ################             ",
		"                        #                                   "})
	if result != 52 {
		t.Fatal(result)
	}
}

func Test3(t *testing.T) {
	result := jumpingGaps([]string{
		"      E   ",
		" # ####   ",
		" #        ",
		"  #####   ",
		"          ",
		"        # ",
		"  S     # ",
		" ######## ",
		"          ",
		"          "})
	if result != 6 {
		t.Fatal(result)
	}
}

func Test1(t *testing.T) {
	result := jumpingGaps([]string{"S#E", "###"})
	if result != 1 {
		t.Fatal(result)
	}
}
