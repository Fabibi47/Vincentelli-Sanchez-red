package player

type Weapon struct {
	Name   string
	Speed  int
	Damage int
	Level  int
	Type   string
	Skills []Skill
}

type Armor struct {
	Helm  string
	Chest string
	Arms  string
	Waist string
	Legs  string
}

var armor Armor

func HpUpdate(armor Armor, player Character) int {
	hpbonus := 0
	hpbonus += HpDetection(armor.Helm)
	hpbonus += HpDetection(armor.Chest)
	hpbonus += HpDetection(armor.Arms)
	hpbonus += HpDetection(armor.Waist)
	hpbonus += HpDetection(armor.Legs)
	return player.Base_hp + hpbonus
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

func SkillDetection(weapon *Weapon) {
	skills := []Skill{}
	if weapon.Type == "GreatSword" {
		skills = append(skills, Slash)
		if weapon.Level >= 3 {
			skills = append(skills, Overhead_Slash)
		}
		if weapon.Level >= 5 {
			skills = append(skills, Charged_Slash)
		}
	}
	if weapon.Type == "LongSword" {
		skills = append(skills, Slash)
		if weapon.Level >= 3 {
			skills = append(skills, Thrust)
		}
		if weapon.Level >= 5 {
			skills = append(skills, Spirit_Slash)
		}
	}
	if weapon.Type == "DualBlades" {
		skills = append(skills, Slash)
		if weapon.Level >= 3 {
			skills = append(skills, Furry_Slash)
		}
		if weapon.Level >= 5 {
			skills = append(skills, Demon_Slash)
		}
	}
	if weapon.Type == "Stick" {
		skills = append(skills, Bonk)
	}
	if weapon.Type == "Sexcalibur" {
		skills = append(skills, Apocalypse)
	}
	weapon.Skills = skills
}
