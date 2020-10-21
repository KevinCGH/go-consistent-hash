package util

import "math"

// 平均值 x=(x1+x2+...+xn)/n
func average(nums []int) float64 {
	m := len(nums)
	if m == 0 {
		return 0
	}
	var sum = 0.0
	for i := 0; i < m; i++ {
		sum += float64(nums[i])
	}
	return sum / float64(m)
}

// 方差 s^2=[(x1-x)^2 + ...+(xn-x)^2]/n
func variance(nums []int) float64 {
	m := len(nums)
	if m == 0 {
		return 0
	}
	var dAve = average(nums)
	var dVar float64 = 0
	for i := 0; i < m; i++ {
		dVar += math.Pow((float64(nums[i]) - dAve), 2)
	}
	return dVar / float64(m)
}

// 标准差 𝞂 = sqrt(s^2)
func standardDeviation(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	dVar := variance(nums)
	return math.Sqrt(dVar)
}
