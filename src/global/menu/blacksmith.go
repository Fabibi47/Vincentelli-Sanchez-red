package menu

import (
	"fmt"
	"player"
	"strconv"
)

func Blacksmith(p player.Character) {
	Clear()
	Write("Welcome! Do you want to forge something ? \n\n\n")
	Write("1 - Weapon \n\n2 - Armor \n\n3 - Leave\n\n")
	selection := ""
	fmt.Scanln(&selection)
	if selection == "1" {
		Weapon()
	} else if selection == "2" {
		Armor(p)
	} else if selection == "3" {
		MainMenu(p)
	} else {
		Write("This is not a valid answer..")
		Blacksmith(p)
	}
}

func Weapon() {
	Clear()

}

func Armor(p player.Character) {

	Clear()
	Write("Armors : \n\n")
	Write("	1 - Great Jagras	\n	2 - Barroth	\n	3 - Rathalos	\n	4 - Nergigante	\n	5 - Xeno'jiiva \n\n	0 - Back\n\n")
	name := ""
	material := ""
	cost := 0
	choice := ""
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		name = "Great Jagras"
		material = "Great Jagras Scale"
		cost = 100
	case "2":
		name = "Barroth"
		material = "Barroth Scalp"
		cost = 200
	case "3":
		name = "Rathalos"
		material = "Rathalos Cortex"
		cost = 350
	case "4":
		name = "Nergigante"
		material = "Nergigante Talon"
		cost = 500
	case "5":
		name = "Xeno'jiiva"
		material = "Xeno'jiiva Gem"
		cost = 750
	case "0":
		Blacksmith(p)
	default:
		Write("not an answer..")
		Armor(p)
	}
	armors := []string{"1 - Helm " + name + " : " + material + " x1 & zenny x" + strconv.Itoa(cost),
		"2 - Chest " + name + " : " + material + " x1 & zenny x" + strconv.Itoa(cost),
		"3 - Arms " + name + " : " + material + " x1 & zenny x" + strconv.Itoa(cost),
		"4 - Waist " + name + " : " + material + " x1 & zenny x" + strconv.Itoa(cost),
		"5 - Legs " + name + " : " + material + " x1 & zenny x" + strconv.Itoa(cost) + "\n",
		"6 - Back",
		"0 - Leave\n"}
	Clear()
	DisplayMenu(armors)
	fmt.Scanln(&choice)
	switch choice {
	case "0":
		MainMenu(p)
	case "6":
		Armor(p)
	}
}
