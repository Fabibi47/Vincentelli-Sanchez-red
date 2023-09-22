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
	armors := []string{"Armors :\n",
		"	1 - Helm " + name + " : " + material + " x1 & zenny x" + strconv.Itoa(cost),
		"	2 - Chest " + name + " : " + material + " x1 & zenny x" + strconv.Itoa(cost),
		"	3 - Arms " + name + " : " + material + " x1 & zenny x" + strconv.Itoa(cost),
		"	4 - Waist " + name + " : " + material + " x1 & zenny x" + strconv.Itoa(cost),
		"	5 - Legs " + name + " : " + material + " x1 & zenny x" + strconv.Itoa(cost) + "\n",
		"6 - Back",
		"0 - Leave\n"}
	DisplayMenu(armors)
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		BuyUpgrade(p, &p.Armor.Helm, &name, material, cost)
	case "2":
		BuyUpgrade(p, &p.Armor.Chest, &name, material, cost)
	case "3":
		BuyUpgrade(p, &p.Armor.Arms, &name, material, cost)
	case "4":
		BuyUpgrade(p, &p.Armor.Waist, &name, material, cost)
	case "5":
		BuyUpgrade(p, &p.Armor.Legs, &name, material, cost)
	case "0":
		MainMenu(p)
	case "6":
		Armor(p)
	}
}

func BuyUpgrade(p *player.Character, part *string, equipment *string, material string, cost int) {
	door := false
	for item, amount := range p.Inventory {
		if item == material && amount >= 1 && p.Money >= cost {
			door = true
		}
	}
	if door == false {
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
			p.Money -= cost
			*part = *equipment
			Write("Thanks brother")
			time.Sleep(3 * time.Second)
			p.Max_health_point = player.HpUpdate(&p.Armor, p)
			Armor(p)
		}
	}
}
