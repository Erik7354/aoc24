package day20

import (
	"slices"
)

type Field = byte

const ( // fields
	Start   Field = 'S'
	End     Field = 'E'
	Wall    Field = '#'
	Track   Field = '.'
	Invalid Field = 0
)

type Racetrack struct {
	Max       Point
	data      [][]Field
	trackpath []Point
}

func NewRacetrack(data [][]Field) Racetrack {
	rt := Racetrack{}
	rt.Max = Point{len(data[0]), len(data)}
	rt.data = data

	s := rt.Find(Start)
	e := rt.Find(End)
	rt.trackpath, _ = rt.Path(s, e)

	return rt
}

func (r Racetrack) At(p Point) Field {
	if (p).In(r) {
		return r.data[p.Y][p.X]
	}
	return Invalid
}

func (r Racetrack) Find(f Field) Point {
	for y := range r.Max.Y {
		for x := range r.Max.X {
			if r.data[y][x] == f {
				return Pt(x, y)
			}
		}
	}

	return Pt(-1, -1)
}

func (r Racetrack) Movable(p Point) bool {
	if !p.In(r) {
		return false
	}

	switch r.data[p.Y][p.X] {
	case Start, End, Track:
		return true
	case Wall:
		return false
	default:
		panic("bug")
	}
}

func (r Racetrack) Distance(start, end Point) (distance int) {
	return slices.Index(r.trackpath, end) - slices.Index(r.trackpath, start)
}

func (r Racetrack) Path(start, end Point) (path []Point, distance int) {
	var visited []Point

	curr := start
	for {
		visited = append(visited, curr)
		distance++

		curr = r.move(curr, visited)
		if curr == end {
			visited = append(visited, end)
			break
		}
	}

	return visited, distance
}

func (r Racetrack) move(curr Point, visited []Point) (next Point) {
	next = curr.Add(North)
	if r.Movable(next) && !slices.Contains(visited, next) {
		return next
	}

	next = curr.Add(East)
	if r.Movable(next) && !slices.Contains(visited, next) {
		return next
	}

	next = curr.Add(South)
	if r.Movable(next) && !slices.Contains(visited, next) {
		return next
	}

	next = curr.Add(West)
	if r.Movable(next) && !slices.Contains(visited, next) {
		return next
	}

	panic("nowhere to go")
}

const costOfCheating = 2

func (r Racetrack) ListCheats(min int) (cheats []Cheat) {
	var next, mustWall Point
	var visited []Point

	for i, curr := range r.trackpath {
		visited = r.trackpath[0:i]

		next = curr.Add(North).Add(North)
		mustWall = curr.Add(North)
		if r.At(mustWall) == Wall && r.Movable(next) && !slices.Contains(visited, next) {
			cheats = append(cheats, Cht(curr, next))
		}

		next = curr.Add(East).Add(East)
		mustWall = curr.Add(East)
		if r.At(mustWall) == Wall && r.Movable(next) && !slices.Contains(visited, next) {
			cheats = append(cheats, Cht(curr, next))
		}

		next = curr.Add(South).Add(South)
		mustWall = curr.Add(South)
		if r.At(mustWall) == Wall && r.Movable(next) && !slices.Contains(visited, next) {
			cheats = append(cheats, Cht(curr, next))
		}

		next = curr.Add(West).Add(West)
		mustWall = curr.Add(West)
		if r.At(mustWall) == Wall && r.Movable(next) && !slices.Contains(visited, next) {
			cheats = append(cheats, Cht(curr, next))
		}
	}

	if min == -1 {
		return cheats
	}

	filtered := make([]Cheat, 0)
	for _, ch := range cheats {
		dist := r.Distance(ch.Start, ch.End)
		if dist-costOfCheating >= min {
			filtered = append(filtered, ch)
		}
	}

	return filtered
}
