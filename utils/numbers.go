package utils

import (

)

func MaxInt64Number(vals...int64) int64 {
	var max int64
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}
func MinInt64Number(vals...int64) int64 {
	var min int64
	for _, val := range vals {
		if min == 0 || val <= min {
			min = val
		}
	}
	return min
}