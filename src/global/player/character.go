package player

type Character struct {
	Name             string
	Level            int
	Max_health_point int
	Health_point     int
	Money            int
	Inventory        map[Item]int
	Armor            Armor
	Weapon           Weapon
	Stamina          int
	Stamina_max      int
	Stack            int
	Affliction       string
}
