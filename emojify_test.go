package main

import "testing"

func TestEmojify(t *testing.T) {
	tests := []struct {
		name  string
		maybe string
		want  string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := emojify(tt.maybe); got != tt.want {
			t.Errorf("%q. emojify() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestMain(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		main()
	}
}

func Benchmark_poop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		emojify(":poop:")

	}

}
