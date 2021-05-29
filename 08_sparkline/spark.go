package main

import (
	"math"
	"strings"
)

// spark returns a sparkline for the provided values.
// This was copied from https://rosettacode.org/wiki/Sparkline_in_unicode#Go.
func spark(vs []float64) (sp string) {
	n := len(vs)
	min := math.Inf(1)
	max := math.Inf(-1)
	for _, v := range vs {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if min == max {
		return strings.Repeat("▄", n)
	}

	rs := make([]rune, n)
	f := 8 / (max - min)
	for j, v := range vs {
		i := rune(f * (v - min))
		if i > 7 {
			i = 7
		}
		rs[j] = '▁' + i
	}
	return string(rs)
}
