package player

type Weapon struct {
	Name   string
	Speed  int
	Damage int
	Level  int
	Type   string
}

type Armor struct {
	Helm  string
	Chest string
	Arms  string
	Waist string
	Legs  string
}

func HpUpdate(armor Armor, player Character) {
	hpbonus := 0
	hpbonus += HpDetection(armor.Helm)
	hpbonus += HpDetection(armor.Chest)
	hpbonus += HpDetection(armor.Arms)
	hpbonus += HpDetection(armor.Waist)
	hpbonus += HpDetection(armor.Legs)
	player.Max_health_point += hpbonus
	player.Health_point += hpbonus
}

func HpDetection(part string) int {
	if part == "Leather" {
		return 0
	} else if part == "Great Jagras" {
		return 5
	} else if part == "Barroth" {
		return 10
	} else if part == "Rathalos" {
		return 20
	} else if part == "Nergigante" {
		return 50
	} else {
		return 100
	}
}
