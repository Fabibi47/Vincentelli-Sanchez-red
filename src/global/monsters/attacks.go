package monsters

type Attacks struct {
	Name   string
	Power  float64
	Effect string
}

func (a Attacks) Use(b Boss) int {
	if a.Effect == "Armor" {
		b.PV = int(float64(b.PV) + a.Power)
		return 0
	}
	return int(float64(b.Damage) * a.Power)
}
