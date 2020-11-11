package math

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		i Inputs
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Add Numbers",
			args: args{i: Inputs{InputA: 3, InputB: 3}},
			want: 6,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.i); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		i Inputs
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Subtract Numbers",
			args: args{i: Inputs{InputA: 4, InputB: 3}},
			want: 1,
		},
		{
			name: "Subtract Numbers",
			args: args{i: Inputs{InputA: 3, InputB: 4}},
			want: -1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.args.i); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
