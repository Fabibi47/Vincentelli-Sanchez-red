package player

type Item struct {
	Name        string
	Descritpion string
	Price       int
	Usable      bool
	Effect      string
	Power       float64
}

func (i Item) Use(p *Character) {
	if p.Inventory[i] > 0 && i.Usable {
		if i.Effect == "Heal" {
			p.Health_point = p.Health_point + int(float64(p.Max_health_point)*i.Power)
			if p.Health_point > p.Max_health_point {
				p.Health_point = p.Max_health_point
			}
			p.Inventory[i]--
		}
	}
}

var First_Aid Item = Item{"First Aid", "Heals 10 percents of you maximum health", 0, true, "Heal", 0.1}
var Potion Item = Item{"Potion", "Heals 25 percents of your maximum health.", 50, true, "Heal", 0.25}
var Mega_Potion Item = Item{"Méga Potion", "Heals 50 percents of your maximum health.", 150, true, "Heal", 0.5}
var Poison Item = Item{"Poison", "Apply on you weapon: gives the effect poison to an ennemy (-5 percents per tour).", 100, true, "Poison", -5.0}
var Jagras_scale = Item{"Jagras Scale", "A scale belonging to a Great Jagras", 50, false, "material", 0.0}
var Barroth_scalp = Item{"Barroth Scalp", "The scalp of a Barroth", 100, false, "material", 0.0}
var Rathalos_cortex = Item{"Rathalos Cortex", "The cortex of a Rathalos", 175, false, "material", 0.0}
var Nergigante_talon = Item{"Nergigante's Talon", "A talon belonging to a Nergigante", 250, false, "material", 0.0}
var Xeno_jiiva_gem = Item{"Xeno' Jiiva's Gem", "A gem belonging to a Xeno' Jiiva", 375, false, "material", 0.0}
var Analtreon = Item{"Analtreon", "Allows you to craft the most powerful weapon and become de King of Hunters", 10000, false, "material", 0.0}
