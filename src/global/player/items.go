package player

type Item struct {
	Name        string
	Descritpion string
	Price       int
	Usable      bool
	Effect      string
	Power       float64
}

func (i Item) Use(p Character) {
	if p.Inventory[i] > 0 && i.Usable {
		if i.Effect == "Heal" {
			p.Health_point += int(float64(p.Max_health_point) * i.Power)
			if p.Health_point > p.Max_health_point {
				p.Health_point = p.Max_health_point
			}
			p.Inventory[i]--
		} else if i.Effect == "Poison" {
			for e := range p.Weapon.Effect {
				delete(p.Weapon.Effect, e)
			}
			p.Weapon.Effect["Poison"] += int(i.Power)
		}
	}
}

var First_Aid Item = Item{"First Aid", "Heals 10 percents of you maximum health", 0, true, "Heal", 1.1}
var Potion Item = Item{"Potion", "Heals 25 percents of your maximum health.", 50, true, "Heal", 1.25}
var Mega_Potion Item = Item{"Méga Potion", "Heals 50 percents of your maximum health.", 150, true, "Heal", 1.5}
var Poison Item = Item{"Poison", "Apply on you weapon: gives the effect poison to an ennemy (-5 percents per tour).", 100, true, "Poison", -5.0}
