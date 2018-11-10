package secretagents

import (
	"fmt"
	"strconv"
	"strings"
)

func secretAgentsMeetingProposal(src string, diff int) []string {
	r, d, m := "", 0, "axxxwyuoietdsnmrbkp"
	a := strings.Split(src, ".")
	for _, s := range a {
		switch s {
		case "*":
			r += "morning"
		case "@":
			r += "afternoon"
		case "#":
			r += "night"
		case "?":
			r += "-"
		case "_":
			d += diff
		default:
			if i, e := strconv.Atoi(s); e == nil {
				fmt.Print(i, d)
				r += string(m[i+d])
				fmt.Println(' ', r)
			} else {
				fmt.Println(e, r)
			}
		}
	}
	a = strings.Split(r, "-")
	ok := false
	switch a[0] {
	case "today":
		ok = a[2] == "park"
	case "tomorrow":
		ok = a[2] == "bar" && a[1] == "night" || a[2] == "park" && a[1] == "afternoon"
	case "twodays":
		ok = a[2] == "restaurant" && a[1] == "morning"
	}
	if ok {
		return []string{r, "5.9.12"}
	} else {
		return []string{r, "13.7"}
	}
}
