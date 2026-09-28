package convexhull

import "testing"

func TestTooFewPoints(t *testing.T) {
	if _, err := Hull([]Point{{X: 0, Y: 0}, {X: 1, Y: 1}}); err == nil {
		t.Fatal("两个点应该报错")
	}
}

func TestTriangle(t *testing.T) {
	hull, err := Hull([]Point{{X: 0, Y: 0}, {X: 2, Y: 0}, {X: 0, Y: 2}})
	if err != nil {
		t.Fatalf("Hull 出错：%v", err)
	}
	if len(hull) != 3 {
		t.Fatalf("三角形凸包应该是 3 个点：%v", hull)
	}
}

func TestSquareCorners(t *testing.T) {
	hull, err := Hull([]Point{{X: 0, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: 2}, {X: 0, Y: 2}})
	if err != nil {
		t.Fatalf("Hull 出错：%v", err)
	}
	if len(hull) != 4 {
		t.Fatalf("正方形凸包应该是 4 个点：%v", hull)
	}
}

func TestInteriorPointIsDropped(t *testing.T) {
	hull, err := Hull([]Point{{X: 0, Y: 0}, {X: 4, Y: 0}, {X: 4, Y: 4}, {X: 0, Y: 4}, {X: 2, Y: 2}})
	if err != nil {
		t.Fatalf("Hull 出错：%v", err)
	}
	for _, point := range hull {
		if point == (Point{X: 2, Y: 2}) {
			t.Fatalf("内部点不该在凸包上：%v", hull)
		}
	}
}

func TestDuplicatePoints(t *testing.T) {
	hull, err := Hull([]Point{{X: 0, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 0}})
	if err != nil {
		t.Fatalf("Hull 出错：%v", err)
	}
	if len(hull) != 3 {
		t.Fatalf("去重后应该是 3 个点：%v", hull)
	}
}
