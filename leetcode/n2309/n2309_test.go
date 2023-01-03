package n2309

import "testing"

func Test_greatestLetter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{
				s: "lEeTcOdE",
			},
			want: "E",
		},
		{
			name: "t2",
			args: args{
				s: "arRAzFif",
			},
			want: "R",
		},
		{
			name: "t3",
			args: args{
				s: "AbCdEfGhIjK",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greatestLetter(tt.args.s); got != tt.want {
				t.Errorf("greatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
