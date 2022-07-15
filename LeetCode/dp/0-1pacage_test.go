package dp

import (
	"reflect"
	"testing"
)

func Test_package0And1(t *testing.T) {
	type args struct {
		N    int
		W    int
		data [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{N: 7, W: 30, data: [][]int{{4, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 9}, {9, 10}, {10, 11}}}, want: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := package0And1(tt.args.N, tt.args.W, tt.args.data); got != tt.want {
				t.Errorf("package0And1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trackback(t *testing.T) {
	type args struct {
		dp   [][]int
		data [][]int
		W    int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{dp: package0And2(7, 21, [][]int{{4, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 9}, {9, 10}, {10, 11}}), data: [][]int{{4, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 9}, {9, 10}, {10, 11}}, W: 21}, want: []bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trackback(tt.args.dp, tt.args.data, tt.args.W); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trackback() = %v, want %v", got, tt.want)
			}
		})
	}
}
