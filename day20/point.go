package day20

var ( // moves
	North = Pt(0, 1)
	East  = Pt(1, 0)
	South = Pt(0, -1)
	West  = Pt(-1, 0)
)

type Point struct {
	X, Y int
}

func Pt(X, Y int) Point {
	return Point{X, Y}
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) In(r Racetrack) bool {
	return 0 <= p.X && p.X < r.Max.X &&
		0 <= p.Y && p.Y < r.Max.Y
}
