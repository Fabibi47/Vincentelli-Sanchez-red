package menu

import (
	"player"
	"strconv"
)

func Stats(p player.Character) {
	stats := []string{
		"Character stats : \n\n\n",
		"	Name : " + p.Name + " | Lvl " + strconv.Itoa(p.Level) + " | Health : " + strconv.Itoa(p.Health_point) + "/" + strconv.Itoa(p.Max_health_point) + " | Zennys : " + strconv.Itoa(p.Money) + "\n\n",
		"	Armor : 			Weapon : \n",
		"	   Helm " + p.Armor.Helm + "			   " + p.Weapon.Name + " " + p.Weapon.Type,
		"	   Chest " + p.Armor.Chest + "		      Damage : " + strconv.Itoa(p.Weapon.Damage),
		"	   Arms " + p.Armor.Arms + "			      Speed : " + strconv.Itoa(p.Weapon.Speed),
		"	   Waist " + p.Armor.Waist,
		"	   Legs " + p.Armor.Legs + "\n",
		"	0 - Back \n\n",
	}
	Clear()
	DisplayMenu(stats)
	scanner.Scan()
	action := scanner.Text()
	switch action {
	default:
		MainMenu(p)
	}
}
