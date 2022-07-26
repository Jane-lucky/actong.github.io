package dp

import "testing"

func Test_putApple1(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{m: 7, n: 3}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putApple1(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("putApple1() = %v, want %v", got, tt.want)
			}
		})
	}
}
