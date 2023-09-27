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

var Slash = Skill{"Slash", 1.0, "A basic slash with your sword dealing 100 percent damage to your target", 0}
var Bonk = Skill{"Bonk", 0.8, "You bonk your ennemy with your stick, dealing 80 percent damange to it", 0}

var Overhead_Slash = Skill{"Overhead Slash", 1.5, "flmm", 25}
var Charged_Slash = Skill{"Charged Slash", 2.0, "flmm", 50}

var Thrust = Skill{"Thrust", 1.5, "flmm", 25}
var Spirit_Slash = Skill{"Spirit Slash", 3.0, "flmm", 75}

var Furry_Slash = Skill{"Furry Slash", 2.0, "flmm", 50}
var Demon_Slash = Skill{"Demon Slash", 5.0, "flmm", 100}

var Apocalypse = Skill{"Apocalypse", 10.0, "...", 0}
