package math

import (
	"testing"
)

func TestInputs_Add(t *testing.T) {
	type fields struct {
		InputA int
		InputB int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Add Numbers",
			fields: fields{InputA: 4, InputB: 3},
			want:   7,
		},
		{
			name:   "Add Numbers",
			fields: fields{InputA: 7, InputB: 25},
			want:   32,
		},
		{
			name:   "Add Numbers",
			fields: fields{InputA: 39, InputB: 3},
			want:   42,
		},
		{
			name:   "Add Numbers",
			fields: fields{InputA: 1, InputB: 2},
			want:   3,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			i := Inputs{
				InputA: tt.fields.InputA,
				InputB: tt.fields.InputB,
			}
			if got := i.Add(); got != tt.want {
				t.Errorf("Inputs.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInputs_Subtract(t *testing.T) {
	type fields struct {
		InputA int
		InputB int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Subtract Numbers",
			fields: fields{InputA: 10, InputB: 3},
			want:   7,
		},
		{
			name:   "Subtract Numbers",
			fields: fields{InputA: 25, InputB: 40},
			want:   -15,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			i := Inputs{
				InputA: tt.fields.InputA,
				InputB: tt.fields.InputB,
			}
			if got := i.Subtract(); got != tt.want {
				t.Errorf("Inputs.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInputs_Multiply(t *testing.T) {
	type fields struct {
		InputA int
		InputB int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Multiply Numbers",
			fields: fields{InputA: 10, InputB: 3},
			want:   30,
		},
		{
			name:   "Multiply Numbers",
			fields: fields{InputA: 25, InputB: 3},
			want:   75,
		},
		{
			name:   "Multiply Numbers",
			fields: fields{InputA: 9, InputB: 3},
			want:   27,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			i := Inputs{
				InputA: tt.fields.InputA,
				InputB: tt.fields.InputB,
			}
			if got := i.Multiply(); got != tt.want {
				t.Errorf("Inputs.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInputs_Divide(t *testing.T) {
	type fields struct {
		InputA int
		InputB int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Divide Numbers",
			fields: fields{InputA: 30, InputB: 3},
			want:   10,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			i := Inputs{
				InputA: tt.fields.InputA,
				InputB: tt.fields.InputB,
			}
			if got := i.Divide(); got != tt.want {
				t.Errorf("Inputs.Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
