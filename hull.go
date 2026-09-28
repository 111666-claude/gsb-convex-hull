// Package convexhull 用单调链求凸包。
// 缺陷：叉积判定用 < 0，共线的中间点会留在结果里。
package convexhull

import (
	"errors"
	"sort"
)

// ErrShape 表示点数不够。
var ErrShape = errors.New("convexhull: 至少需要 3 个点")

// Point 是一个平面点。
type Point struct {
	X float64
	Y float64
}

func cross(origin, first, second Point) float64 {
	return (first.X-origin.X)*(second.Y-origin.Y) - (first.Y-origin.Y)*(second.X-origin.X)
}

// Hull 返回逆时针凸包。缺陷：共线点不会被丢掉。
func Hull(points []Point) ([]Point, error) {
	if len(points) < 3 {
		return nil, ErrShape
	}
	sorted := append([]Point(nil), points...)
	sort.Slice(sorted, func(left, right int) bool {
		if sorted[left].X != sorted[right].X {
			return sorted[left].X < sorted[right].X
		}
		return sorted[left].Y < sorted[right].Y
	})
	unique := sorted[:1]
	for _, point := range sorted[1:] {
		if point != unique[len(unique)-1] {
			unique = append(unique, point)
		}
	}
	if len(unique) < 3 {
		return nil, ErrShape
	}
	build := func(chain []Point, point Point) []Point {
		for len(chain) >= 2 && cross(chain[len(chain)-2], chain[len(chain)-1], point) < 0 {
			chain = chain[:len(chain)-1]
		}
		return append(chain, point)
	}
	lower := []Point{}
	for _, point := range unique {
		lower = build(lower, point)
	}
	upper := []Point{}
	for index := len(unique) - 1; index >= 0; index-- {
		upper = build(upper, unique[index])
	}
	return append(lower[:len(lower)-1], upper[:len(upper)-1]...), nil
}
