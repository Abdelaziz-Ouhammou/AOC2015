package main

import "testing"

func TestFindFloor(t *testing.T) {
	cases := map[string]int{
		"(())":    0,
		"()()":    0,
		"(((":     3,
		"(()(()(": 3,
		"))(((((": 3,
		"())":     -1,
		"))(":     -1,
		")))":     -3,
		")())())": -3,
	}
	for c, f := range cases {
		r := findFloor(c)
		if r != f {
			t.Errorf("expected: %d got: %d instead\n", f, r)
		}
	}
}

func TestReadInput(t *testing.T) {
	str, err := readInput("testdata/input.txt")
	if err != nil {
		t.Error("failed to read input file")
	}
	if str != "((((()))))" {
		t.Errorf("expected: %s got %s instead", "((((()))))", str)
	}
}

func TestFindFirstBasementEntry(t *testing.T) {
	cases := map[string]int{
		")":     1,
		"()())": 5,
		"())":   3,
		"(((":   -1,
	}
	for c := range cases {
		instructionNumber := findFirstBasementEntry(c)
		if instructionNumber != cases[c] {
			t.Errorf("expected: %d got %d instead", cases[c], instructionNumber)
		}
	}
}
