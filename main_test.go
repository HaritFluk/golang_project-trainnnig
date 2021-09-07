package main

import "testing"

func TestPlus(t *testing.T)  {
	a := plus(1, 2)

	if a != 3 {
		t.Errorf("Expected result of J, but it was %v")
	}


}