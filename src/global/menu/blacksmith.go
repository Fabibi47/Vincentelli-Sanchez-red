package menu

import (
	"fmt"
	"player"
	"strconv"
	"time"
)

func Blacksmith(p *player.Character) {
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

func Armor(p *player.Character) {

	Clear()
	Write("Armors : \n\n")
	Write("	1 - Great Jagras (5 hp per part)	\n	2 - Barroth (10 hp per part)	\n	3 - Rathalos (20 hp per part)	\n	4 - Nergigante (50 hp per part)	\n	5 - Xeno'jiiva (100 hp per part) \n\n	0 - Back\n\n")
	var material player.Item
	name := ""
	choice := ""
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		name = "Great Jagras"
		material = player.Jagras_scale
	case "2":
		name = "Barroth"
		material = player.Barroth_scalp
	case "3":
		name = "Rathalos"
		material = player.Rathalos_cortex
	case "4":
		name = "Nergigante"
		material = player.Nergigante_talon
	case "5":
		name = "Xeno' Jiiva"
		material = player.Xeno_jiiva_gem
	case "0":
		Blacksmith(p)
	default:
		Write("not an answer..")
		Armor(p)
	}
	armors := []string{"Armors :\n",
		"	1 - Helm " + name + " : " + material.Name + " x1 & zenny x" + strconv.Itoa(material.Price*2),
		"	2 - Chest " + name + " : " + material.Name + " x1 & zenny x" + strconv.Itoa(material.Price*2),
		"	3 - Arms " + name + " : " + material.Name + " x1 & zenny x" + strconv.Itoa(material.Price*2),
		"	4 - Waist " + name + " : " + material.Name + " x1 & zenny x" + strconv.Itoa(material.Price*2),
		"	5 - Legs " + name + " : " + material.Name + " x1 & zenny x" + strconv.Itoa(material.Price*2) + "\n",
		"6 - Back",
		"0 - Leave\n"}
	Clear()
	DisplayMenu(armors)
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		BuyUpgrade(p, &p.Armor.Helm, &name, material)
	case "2":
		BuyUpgrade(p, &p.Armor.Chest, &name, material)
	case "3":
		BuyUpgrade(p, &p.Armor.Arms, &name, material)
	case "4":
		BuyUpgrade(p, &p.Armor.Waist, &name, material)
	case "5":
		BuyUpgrade(p, &p.Armor.Legs, &name, material)
	case "0":
		MainMenu(p)
	case "6":
		Armor(p)
	}
}

func BuyUpgrade(p *player.Character, part *string, equipment *string, material player.Item) {
	door := false
	for item, amount := range p.Inventory {
		if item == material && amount >= 1 && p.Money >= material.Price*2 {
			door = true
		}
	}
	if !door {
		Write("You don't have the materials required brother")
		time.Sleep(3 * time.Second)
		Armor(p)
	} else {
		if part == equipment {
			Write("You already have that equipped")
			time.Sleep(3 * time.Second)
			Armor(p)
		} else {
			p.Inventory[material] -= 1
			p.Money -= material.Price * 2
			*part = *equipment
			Write("Thanks brother")
			time.Sleep(3 * time.Second)
			p.Max_health_point = player.HpUpdate(&p.Armor, p)
			Armor(p)
		}
	}
}
