package player

type Skill struct {
	Name        string
	Damage      float64
	Description string
	Cost        int
}

func (s Skill) Use(p *Character) int {
	damage := p.Weapon.Damage * int(s.Damage)
	return damage
}

var Slash = Skill{"Slash", 1.0, "A basic slash with your sword dealing 100 percent damage to your target", 10}
var Bonk = Skill{"Bonk", 0.8, "You bonk your ennemy with your stick, dealing 80 percent damange to it", 5}
