package player

type Character struct {
	Name             string
	Level            int
	Max_health_point int
	Health_point     int
	Money            int
	Inventory        map[string]int
	Armor            Armor
	Weapon           Weapon
}
