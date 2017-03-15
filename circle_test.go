package main

import (
	"math"
	"testing"
)

func testCircleIntersection(t *testing.T) {
	c1 := Circle{X: -3, Y: 0, Radius: 6}
	c2 := Circle{X: 3, Y: 0, Radius: 6}
	area := circleIntersection(c1, c2)

	expect_area := 24*math.Pi - 18*math.Sqrt(3)
	if area != expect_area {
		t.Fatal("expect: %s, got: %s", expect_area, area)
	}
}
