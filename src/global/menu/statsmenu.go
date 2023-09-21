package menu

import (
	"player"
	"strconv"
)

func Stats(p player.Character) {
	stats := []string{"Name :" + p.Name + "\n",
		"Weapon :" + p.Weapon.Name + "\n",
		"	Speed :" + strconv.Itoa(p.Weapon.Speed) + "\n",
		"	Damage :" + strconv.Itoa(p.Weapon.Damage) + "\n",
		"Level :" + strconv.Itoa(p.Level) + "\n",
		"Health :" + strconv.Itoa(p.Health_point) + "/" + strconv.Itoa(p.Max_health_point) + "\n",
		"Money :" + strconv.Itoa(p.Money) + "\n",
		"\n0 - Back",
	}
	DisplayMenu(stats)
	scanner.Scan()
	action := scanner.Text()
	switch action {
	default:
		MainMenu(p)
	}
}
