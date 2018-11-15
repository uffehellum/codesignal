package clockhands

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

func clockHandAngle(time string) (x float64) {
	a := strings.Split(time, ":")
	for _, t := range a {
		p, _ := strconv.ParseFloat(t, 64)
		x = x*60 + p
	}
	fmt.Println(x)
	x = x / 10.0 * 59.0
	fmt.Println(x)
	x = math.Remainder(x-180, 360) + 180
	fmt.Println(x)
	x = math.Abs(x)
	fmt.Println(x)
	x *= 360
	return
}

func TestRemainder(t *testing.T) {
	for i := 0.0; i < 50; i++ {
		fmt.Println(i, math.Remainder(i, 20))
	}
	t.Fail()
}
