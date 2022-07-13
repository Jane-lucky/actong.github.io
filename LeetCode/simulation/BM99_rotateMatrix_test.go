package simulation

import (
	"reflect"
	"testing"
)

func Test_rotateMatrix(t *testing.T) {
	type args struct {
		mat [][]int
		n   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, n: 3}, want: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateMatrix(tt.args.mat, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
