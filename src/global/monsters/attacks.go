package monsters

type Attacks struct {
	Name   string
	Power  float64
	Effect string
}

func (a Attacks) Use(b *Boss) int {
	if a.Effect == "Armor" {
		b.PV += int(a.Power)
		return 0
	} else {
		return b.Damage * int(a.Power)
	}
}
