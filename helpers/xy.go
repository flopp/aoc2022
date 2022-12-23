package helpers

type XY struct {
	X, Y int
}

func (xy XY) Offset(x, y int) XY {
	return XY{xy.X + x, xy.Y + y}
}

func (xy XY) Manhattan(other XY) int {
	dx := xy.X - other.X
	if dx < 0 {
		dx = -dx
	}

	dy := xy.X - other.Y
	if dy < 0 {
		dy = -dy
	}

	return dx + dy
}
