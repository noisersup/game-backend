package models

type Player struct {
	Name  string
	Loc   Location
	Speed float64
}

func InitPlayer(name string, x, y, speed float64) Player {
	return Player{Name: name, Loc: Location{X: x, Y: y}, Speed: speed}
}
func (p *Player) SetLocation(loc Location) {
	p.Loc = loc
}
func (p *Player) MoveForward() {
	p.SetLocation(Location{X: p.Loc.X + p.Speed, Y: p.Loc.Y + p.Speed})
}

type Location struct {
	X float64
	Y float64
}

type Map struct {
	XSize float64
	YSize float64
}
