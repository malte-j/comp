package utils

import "testing"

func Test_FirstLetterUppercase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first letter uppercase",
			args: args{
				s: "hello",
			},
			want: "Hello",
		},
		{
			name: "empty string",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "one letter",
			args: args{
				s: "a",
			},
			want: "A",
		},
		{
			name: "numbers",
			args: args{
				s: "123",
			},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstLetterUppercase(tt.args.s); got != tt.want {
				t.Errorf("firstLetterUppercase() = %v, want %v", got, tt.want)
			}
		})
	}
}
