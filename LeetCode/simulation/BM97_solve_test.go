package simulation

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		m int
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{n: 6, m: 2, a: []int{1, 2, 3, 4, 5, 6}}, want: []int{5, 6, 1, 2, 3, 4}},
		{name: "test2", args: args{n: 6, m: 7, a: []int{1, 2, 3, 4, 5, 6}}, want: []int{6, 1, 2, 3, 4, 5}},
		{name: "test2", args: args{n: 1, m: 5, a: []int{5}}, want: []int{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.m, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
