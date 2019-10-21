package main

import (
	"reflect"
	"testing"
)

func Benchmark_generateRandomText(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		generateRandomText(1_000_000)
	}
}

func Test_generateRandomText(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{
			name:    "10",
			args:    args{10},
			wantLen: 10,
		},
		{
			name:    "1kk",
			args:    args{1_000_000},
			wantLen: 1_000_000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotText := generateRandomText(tt.args.n); !reflect.DeepEqual(len(gotText), tt.wantLen) {
				t.Errorf("generateRandomText() = %v, want %v", len(gotText), tt.wantLen)
			}
		})
	}
}
