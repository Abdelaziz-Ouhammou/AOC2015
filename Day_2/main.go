package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rect struct {
	l int
	w int
	h int
}

func main() {
	// Part 1
	rects, err := parseRectsFromFile("input.txt")
	if err != nil {
		panic(err)
	}
	var sumOfPaper int
	for _, r := range rects {
		sumOfPaper += calcPaperRequired(r)
	}
	fmt.Println("total amount of paper required is:", sumOfPaper)
	// Part 2
	var sumOfRibon int
	for _, r := range rects {
		sumOfRibon += calcRibonRequired(r)
	}
	fmt.Println("total length of ribon required is:", sumOfRibon)
}

func parseRect(input string) (Rect, error) {
	dimsString := strings.Split(input, "x")
	var err error
	r := Rect{}
	r.l, err = strconv.Atoi(dimsString[0])
	if err != nil {
		return Rect{}, err
	}
	r.w, err = strconv.Atoi(dimsString[1])
	if err != nil {
		return Rect{}, err
	}
	r.h, err = strconv.Atoi(dimsString[2])
	if err != nil {
		return Rect{}, err
	}
	return r, nil
}
func parseRectsFromFile(filename string) ([]Rect, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []Rect{}, err
	}
	s := bufio.NewScanner(f)
	rects := []Rect{}
	for s.Scan() {
		r, err := parseRect(s.Text())
		if err != nil {
			return []Rect{}, err
		}
		rects = append(rects, r)
	}
	return rects, nil
}

func calcPaperRequired(r Rect) int {
	sides := []int{r.l * r.w, r.l * r.h, r.h * r.w}
	smallSide := sides[0]
	sum := 0
	for i := range sides {
		if sides[i] < smallSide {
			smallSide = sides[i]
		}
		sum += sides[i] * 2
	}
	return sum + smallSide
}

func calcRibonRequired(r Rect) int {
	dims := []int{r.l, r.w, r.h}
	largestDim := dims[0]

	sum := 0
	for i := range dims {
		if dims[i] > largestDim {
			largestDim = dims[i]
		}
		sum += dims[i] * 2
	}
	volume := r.l * r.w * r.h
	sum += volume - (largestDim * 2)
	return sum
}
