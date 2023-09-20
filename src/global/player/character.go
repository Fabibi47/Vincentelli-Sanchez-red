package player

type Character struct {
	Name             string
	Weapon           string
	Speed            int
	Damage           int
	Level            int
	Max_health_point int
	Health_point     int
	Money            int
	Inventory        map[string]int
}
