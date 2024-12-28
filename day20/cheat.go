package day20

type Cheat struct {
	Start, End Point
}

func Cht(s, e Point) Cheat {
	return Cheat{s, e}
}
