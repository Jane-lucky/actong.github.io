package Alth

import "testing"

func TestMaxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, want: 49},
		{name: "2", args: args{[]int{1, 1}}, want: 1},
		{name: "3", args: args{[]int{}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxArea(tt.args.height); got != tt.want {
				t.Errorf("MaxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}