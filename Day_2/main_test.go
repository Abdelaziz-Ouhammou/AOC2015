package main

import "testing"

func TestParseRectangle(t *testing.T) {
	input := "1x2x3"
	r := Rect{l: 1, w: 2, h: 3}
	result, err := parseRect(input)
	if err != nil {
		t.Errorf("failed to parse Rect: %s", err)
	}
	if r != result {
		t.Errorf("expected %#v, got instead %#v", r, result)
	}
	input = "bad input"
	_, err = parseRect(input)
	if err == nil {
		t.Errorf("should return error for invalid input")
	}
}

func TestParseRectsFromFile(t *testing.T) {
	r := Rect{1, 2, 3}
	rects, err := parseRectsFromFile("testdata/input.txt")
	if err != nil {
		t.Errorf("could not read the input file: %#v", err)
	}
	if r != rects[0] {
		t.Errorf("expected: %#v, got instead %#v", r, rects[0])
	}
}

func TestCalcPaperRequired(t *testing.T) {
	cases := map[Rect]int{
		{2, 3, 4}:  58,
		{1, 1, 10}: 43,
	}
	for rect, surface := range cases {
		result := calcPaperRequired(rect)
		if surface != result {
			t.Errorf("%#v: expected: %#v, got instead %#v", rect, surface, result)
		}
	}
}

func TestCalcRibonRequired(t *testing.T) {
	cases := map[Rect]int{
		{2, 3, 4}:  34,
		{1, 1, 10}: 14,
	}
	for rect, length := range cases {
		result := calcRibonRequired(rect)
		if length != result {
			t.Errorf("%#v: expected: %#v, got instead %#v", rect, length, result)
		}
	}
}
