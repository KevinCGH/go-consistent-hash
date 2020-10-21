package util

import (
	"math"
	"testing"
)

type want struct {
	average           float64
	variance          float64
	standardDeviation float64
}

var tests = []struct {
	name string
	args []int
	want want
}{
	// TODO: Add test cases.
	{"empty", []int{}, want{0.0, 0.0, 0.0}},
	{"case1", []int{1, 2, 3, 4, 5}, want{3, 2, math.Sqrt(2)}},
}

func Test_average(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.args); got != tt.want.average {
				t.Errorf("average() = %v, want %v", got, tt.want.average)
			}
		})
	}
}

func Test_variance(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := variance(tt.args); got != tt.want.variance {
				t.Errorf("variance() = %v, want %v", got, tt.want.variance)
			}
		})
	}
}

func Test_standardDeviation(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := standardDeviation(tt.args); got != tt.want.standardDeviation {
				t.Errorf("standardDeviation() = %v, want %v", got, tt.want.standardDeviation)
			}
		})
	}
}
