package main

import "testing"

func Test_fizzBuzzNormal(t *testing.T) {

	tests := []struct {
		name string
		args int
		want string
	}{
		{
			name: "case fizz",
			args: 3,
			want: "Fizz",
		}, {
			name: "case buzz",
			args: 20,
			want: "Buzz",
		}, {
			name: "case fizzbuzz",
			args: 30,
			want: "FizzBuzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzzNormal(tt.args); got != tt.want {
				t.Errorf("fizzBuzzNormal() = %v, want %v", got, tt.want)
			}
		})
	}
}
