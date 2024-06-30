package main

import (
	"testing"
)

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
		{
			name: "case number",
			args: 2,
			want: "2",
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

func Test_fizbuzz1if(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case fizz",
			args: args{3},
			want: "Fizz",
		}, {
			name: "case buzz",
			args: args{5},
			want: "Buzz",
		}, {
			name: "case fizzbuzz",
			args: args{60},
			want: "FizzBuzz",
		},
		{
			name: "case number",
			args: args{2},
			want: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizbuzz1if(tt.args.n); got != tt.want {
				t.Errorf("fizbuzz1if() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fizznoif(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case fizz",
			args: args{3},
			want: "Fizz",
		}, {
			name: "case buzz",
			args: args{5},
			want: "Buzz",
		}, {
			name: "case fizzbuzz",
			args: args{60},
			want: "FizzBuzz",
		},
		{
			name: "case number",
			args: args{2},
			want: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizznoif(tt.args.n); got != tt.want {
				t.Errorf("fizznoif() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_romantonumber(t *testing.T) {

	tests := []struct {
		name string
		args string
		want int
	}{
		{name: "case 1", args: "I", want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToNumber(tt.args); got != tt.want {
				t.Errorf("romantonumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
