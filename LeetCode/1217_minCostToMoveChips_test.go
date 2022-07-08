package Alth

import "testing"

func Test_minCostToMoveChips(t *testing.T) {
	type args struct {
		position []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{position: []int{1, 2, 3}}, want: 1},
		{name: "test2", args: args{position: []int{2, 2, 2, 3, 3}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostToMoveChips(tt.args.position); got != tt.want {
				t.Errorf("minCostToMoveChips() = %v, want %v", got, tt.want)
			}
		})
	}
}
