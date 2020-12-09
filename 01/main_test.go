package main

import "testing"

var sampleInput = []int{
	1721,
	979,
	366,
	299,
	675,
	1456,
}

func TestFindTwoWithSum(t *testing.T) {
	expected := 514579

	if got := FindTwoWithSum(sampleInput, 2020); got != expected {
		t.Errorf("Expected: %d; Got: %d", expected, got)
	}
}

func TestFindThreeWithSum(t *testing.T) {
	expected := 514579

	if got := FindTwoWithSum(sampleInput, 2020); got != expected {
		t.Errorf("Expected: %d; Got: %d", expected, got)
	}
}
