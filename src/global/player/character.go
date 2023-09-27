package player

type Character struct {
	Name             string
	Level            int
	Exp              int
	Base_hp          int
	Max_health_point int
	Health_point     int
	Money            int
	Inventory        map[Item]int
	Armor            Armor
	Weapon           Weapon
	Stamina          int
	Stamina_max      int
<<<<<<< HEAD
	Stack            int
	Affliction       string
=======
}

func LevelUpdate(p *Character) {
	upgrading := true
	for upgrading {
		if p.Exp >= p.Level*10 {
			p.Exp -= p.Level * 10
			p.Level++
			p.Base_hp += 5
			p.Max_health_point = HpUpdate(&p.Armor, p)
		} else {
			upgrading = false
		}
	}

>>>>>>> f1860076939a29938ee7eebbced9925ae3096641
}
