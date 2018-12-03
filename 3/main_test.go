package main

import (
	"strings"
	"testing"
)

func TestPartA(t *testing.T) {
	testdata := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`
	wanted := 4
	got := partA(strings.Split(testdata, "\n"), 10)
	if got != wanted {
		t.Errorf("Wanted %v but got %v", wanted, got)
	}
}

func TestPartB(t *testing.T) {
	testdata := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`
	wanted := 3
	got := partB(strings.Split(testdata, "\n"))
	if got != wanted {
		t.Errorf("Wanted %v but got %v", wanted, got)
	}
}

func TestPartB2(t *testing.T) {
	testdata := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`
	wanted := 3
	got := partB2(strings.Split(testdata, "\n"), 10)
	if got != wanted {
		t.Errorf("Wanted %v but got %v", wanted, got)
	}
}

func BenchmarkPartA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partA(strings.Split(data, "\n"), 1000)
	}
}

func BenchmarkPartB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partB(strings.Split(data, "\n"))
	}
}

func BenchmarkPartB2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partB2(strings.Split(data, "\n"), 1000)
	}
}
