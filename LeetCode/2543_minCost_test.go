package Alth

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {name: "test1", args: args{costs: [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}}, want: 10},
		{name: "test2", args: args{costs: [][]int{{3, 5, 3}, {6, 17, 6}, {7, 13, 18}, {9, 10, 18}}}, want: 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.costs); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
