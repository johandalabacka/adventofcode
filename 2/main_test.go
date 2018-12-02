package main

import (
	"testing"
)

func TestPartA(t *testing.T) {
	wanted := 8398
	got := partA()
	if got != wanted {
		t.Errorf("Wanted %v but got %v", wanted, got)
	}
}

func TestPartB(t *testing.T) {
	wanted := "hhvsdkatysmiqjxunezgwcdpr"
	got := partB()
	if got != wanted {
		t.Errorf("Wanted %v but got %v", wanted, got)
	}
}

func TestPartB2(t *testing.T) {
	wanted := "hhvsdkatysmiqjxunezgwcdpr"
	got := partB2()
	if got != wanted {
		t.Errorf("Wanted %v but got %v", wanted, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partA()
	}
}

func BenchmarkPartB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partB()
	}
}

func BenchmarkPartB2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partB2()
	}
}
