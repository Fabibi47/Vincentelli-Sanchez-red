package menu

import (
	"red/global/player"
	"strconv"
)

func Stats(p *player.Character) {
	stats := []string{
		"" + "\033[36m" + " .----..---.  .-" + "\033[34m" + "-.  .---.  .----.   " + "\033[0m" + "",
		"" + "\033[36m" + "{ {__ {_   _}/ {} \\{_   _}{ " + "\033[34m" + "{__      _/" + "\033[0m" + "",
		"" + "\033[36m" + ".-._} } | | /  /" + "\033[34m" + "\\  \\ | |  .-._} }   " + "\033[0m" + "",
		"" + "\033[36m" + "`----'  `-' `-'  `-' `" + "\033[34m" + "-'  `----'     _/   \n\n" + "\033[0m" + "",
		"	Name : " + "\033[90m" + "" + p.Name + "" + "\033[0m" + " | " + "\033[92m" + "Lvl " + strconv.Itoa(p.Level) + "" + "\033[0m" + " | " + "\033[91m" + "Health : " + strconv.Itoa(p.Health_point) + "/" + strconv.Itoa(p.Max_health_point) + "" + "\033[0m" + " |" + "\033[93m" + " Zennys : " + strconv.Itoa(p.Money) + "" + "\033[0m" + "\n",
		"	Armor :\n" + "\033[37m" + "",
		"	   Helm " + p.Armor.Helm,
		"	   Chest " + p.Armor.Chest,
		"	   Arms " + p.Armor.Arms,
		"	   Waist " + p.Armor.Waist,
		"	   Legs " + p.Armor.Legs + "\n" + "\033[0m" + "",
		"	Weapon :\n" + "\033[37m" + "",
		"	   " + p.Weapon.Name + " " + p.Weapon.Type,
		"" + "\033[0m" + "	   Damage : " + "\033[31m" + "" + strconv.Itoa(p.Weapon.Damage) + "\033[0m",
		"" + "\033[0m" + "	   Speed : " + "\033[35m" + "" + strconv.Itoa(p.Weapon.Speed) + "" + "\033[0m" + "\n",
		"" + "\033[0m" + "	   Skills :\n",
	}
	for _, s := range p.Weapon.Skills {
		stats = append(stats, "	     "+s.Name+" - Damage : "+"\033[31m"+""+strconv.Itoa(int(s.Damage*float64(p.Weapon.Damage)))+""+"\033[0m"+" | Endurance :"+"\033[36m"+" "+strconv.Itoa(s.Cost))
	}
	stats = append(stats, ""+"\033[0m"+"\n0 - Back\n")
	Clear()
	DisplayMenu(stats)
	scanner.Scan()
	action := scanner.Text()
	switch action {
	default:
		MainMenu(p)
	}
}
