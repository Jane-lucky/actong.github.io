package dp

import "testing"

func Test_plan1(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{n: 1, m: 2}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plan1(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("plan() = %v, want %v", got, tt.want)
			}
		})
	}
}
