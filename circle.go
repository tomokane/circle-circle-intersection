package main

import (
	"math"
)

type Circle struct {
	X      int
	Y      int
	Radius float64
}

func abs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}

// 二つの円の重なる部分の面積を求める
func circleIntersection(c1 Circle, c2 Circle) float64 {
	// 二つの円の中心点の距離
	d := distance(c1, c2)

	d1 := (math.Pow(d, 2) - math.Pow(c2.Radius, 2) + math.Pow(c1.Radius, 2)) / (2 * d)
	d2 := d - d1

	// 重なってる部分の面積
	return circularSegment(c1.Radius, d1) + circularSegment(c2.Radius, d2)
}

// 二つの円の中心点の距離
func distance(c1 Circle, c2 Circle) float64 {
	dx := float64(abs(c1.X - c2.X))
	dy := float64(abs(c1.Y - c2.Y))
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

// Circular Segmentの面積を計算
// rは円の半径、dは円の中心点から底辺の直角で接する距離
func circularSegment(r float64, d float64) float64 {
	a := math.Pow(r, 2) * math.Acos(d/r)
	b := d * math.Sqrt(math.Pow(r, 2)-math.Pow(d, 2))
	return a - b
}
