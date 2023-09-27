package menu

import (
	"player"
	"strconv"
)

func Stats(p *player.Character) {
	stats := []string{
		"Character stats : \n\n\n",
		"	Name : " + p.Name + " | Lvl " + strconv.Itoa(p.Level) + " | Health : " + strconv.Itoa(p.Health_point) + "/" + strconv.Itoa(p.Max_health_point) + " | Stamina : " + strconv.Itoa(p.Stamina) + "/" + strconv.Itoa(p.Stamina_max) + " | Zennys : " + strconv.Itoa(p.Money) + "\n\n",
		"	Armor :\n",
		"	   Helm " + p.Armor.Helm,
		"	   Chest " + p.Armor.Chest,
		"	   Arms " + p.Armor.Arms,
		"	   Waist " + p.Armor.Waist,
		"	   Legs " + p.Armor.Legs + "\n",
		"	Weapon :\n",
		"	   " + p.Weapon.Name + " " + p.Weapon.Type,
		"	   Damage : " + strconv.Itoa(p.Weapon.Damage),
		"	   Speed : " + strconv.Itoa(p.Weapon.Speed) + "\n",
		"	   Skills :\n",
	}
	for _, s := range p.Weapon.Skills {
		stats = append(stats, "	     "+s.Name)
	}
	stats = append(stats, "\n0 - Back\n")
	Clear()
	DisplayMenu(stats)
	scanner.Scan()
	action := scanner.Text()
	switch action {
	default:
		MainMenu(p)
	}
}
