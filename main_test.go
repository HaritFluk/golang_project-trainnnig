package main

import "testing"

func TestPlus(t *testing.T)  {
	a := plus(1, 2)

	if a != 3 {
		t.Errorf("Expected result of J, but it was %v instead",a)
	}
}

func TestPrism(t *testing.T) {
	a := prism(4, 2, 2)

	if a != 16 {
		t.Errorf("Expected result of 16, but it was %v instead",a)
	}
}

func BenchmarkFact10(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		fact(10)
	}
}

func BenchmarkFactUsingFor10(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		factUsingFor(10)
	}
}