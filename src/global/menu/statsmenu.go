package menu

import (
	"fmt"
	"player"
	"strconv"
)

func Stats(p player.Character) {
	stats := []string{"Name :" + p.Name + "\n",
		"Weapon :" + p.Weapon.Name + "\n",
		"	Speed :" + strconv.Itoa(p.Weapon.Speed) + "\n",
		"	Damage :" + strconv.Itoa(p.Weapon.Damage) + "\n",
		"Level :" + strconv.Itoa(p.Level) + "\n",
		"Health :" + fmt.Sprintf("%v", p.Health_point) + "/" + fmt.Sprintf("%v", p.Health_point) + "\n",
		"Money :" + strconv.Itoa(p.Money) + " Zennys \n",
		"\n0 - Back",
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
