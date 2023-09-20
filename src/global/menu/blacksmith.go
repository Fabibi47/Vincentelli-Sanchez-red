package menu

import (
	"fmt"
)

func Blacksmith() {
	Clear()
	Write("Welcome! Do you want to forge items ? \n\n\n")
	Write("1 - Weapon \n\n2 - Armor \n\n3 - Leave")
	selection := ""
	fmt.Scanln(&selection)
	if selection == "1" {
		Weapon()
	} else if selection == "2" {
		Armor()
	} else if selection == "3" {
		MainMenu()
	} else {
		Write("This is not a valid answer..")
		Blacksmith()
	}
}

func Weapon() {
	Clear()

}

func Armor() {

	Clear()
	Write("Armors : \n\n")
	Write("1 - Great Jagras	\n	2 - Barroth	\n	3 - Rathalos	\n	4 - Nergigante	\n	5 - Xeno'jiiva \n0 - Back")
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
		Blacksmith()
	default:
		Write("not an answer..")
		Armor()
	}
	Write("1 - Helms " + name + " : " + material + " x1 & zenny x" + string(cost) + "\n")
	Write("2 - Chests " + name + " : " + material + " x1 & zenny x" + string(cost) + "\n")
	Write("3 - Arms " + name + " : " + material + " x1 & zenny x" + string(cost) + "\n")
	Write("4 - Waist " + name + " : " + material + " x1 & zenny x" + string(cost) + "\n")
	Write("5 - Legs " + name + " : " + material + " x1 & zenny x" + string(cost) + "\n")
	Write("6 - Back \n")
	Write("0 - Leave")
	fmt.Scanln(&choice)
	switch choice {
	case "0":
		MainMenu()
	case "6":
		Armor()
	}
}
