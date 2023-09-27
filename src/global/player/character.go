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

}
