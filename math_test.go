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
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		i Inputs
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Multiply Numbers",
			args: args{i: Inputs{InputA: 4, InputB: 3}},
			want: 12,
		},
		{
			name: "Multiply Numbers",
			args: args{i: Inputs{InputA: 3, InputB: 5}},
			want: 15,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.i); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		i Inputs
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Divide Numbers",
			args: args{i: Inputs{InputA: 12, InputB: 3}},
			want: 4,
		},
		{
			name: "Divide Numbers",
			args: args{i: Inputs{InputA: 15, InputB: 5}},
			want: 3,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.args.i); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
