package menu

import (
	"fmt"
	"red/global/player"

	"strconv"
	"time"
)

func Blacksmith(p *player.Character) {
	Clear()
	blacksmith := []string{
		"\033[95m" + "__________.__                 __                   .__  __  .__          ",
		"\\______   \\  | _____    ____ |  | __  ______ _____ |__|/  |_|  |__    /\\ ",
		" |    |  _/  | \\__  \\ _/ ___\\|  |/ / /  ___//     \\|  \\   __\\  |  \\   \\/ ",
		" |    |   \\  |__/ __ \\\\  \\___|    <  \\___ \\|  Y Y  \\  ||  | |   Y  \\  /\\ ",
		" |______  /____(____  /\\___  >__|_ \\/____  >__|_|  /__||__| |___|  /  \\/ ",
		"        \\/          \\/     \\/     \\/     \\/      \\/              \\/     \n\n" + "\033[0m",
	}
	DisplayMenu(blacksmith)
	Write("Welcome! Do you want to forge something ? \n\n\n")
	Write("1 - Weapon \n\n2 - Armor \n\n0 - Leave\n\n")
	selection := ""
	fmt.Scanln(&selection)
	if selection == "1" {
		Weapon(p)
	} else if selection == "2" {
		Armor(p)
	} else if selection == "0" {
		MainMenu(p)
	} else {
		Write("This is not a valid answer..")
		Blacksmith(p)
	}
}

func Weapon(p *player.Character) {
	Clear()
	blacksmith := []string{
		"\033[95m" + "__________.__                 __                   .__  __  .__          ",
		"\\______   \\  | _____    ____ |  | __  ______ _____ |__|/  |_|  |__    /\\ ",
		" |    |  _/  | \\__  \\ _/ ___\\|  |/ / /  ___//     \\|  \\   __\\  |  \\   \\/ ",
		" |    |   \\  |__/ __ \\\\  \\___|    <  \\___ \\|  Y Y  \\  ||  | |   Y  \\  /\\ ",
		" |______  /____(____  /\\___  >__|_ \\/____  >__|_|  /__||__| |___|  /  \\/ ",
		"        \\/          \\/     \\/     \\/     \\/      \\/              \\/     \n\n" + "\033[0m",
	}
	DisplayMenu(blacksmith)
	Write("Weapons : \n\n")
	Write("	1 - Buy Weapons \n	2 - Upgrade Weapons \n\n0 - Back\n\n")
	choice := ""
	fmt.Scanln(&choice)
	Clear()
	switch choice {
	case "1":
		blacksmith := []string{
			"\033[95m" + "__________.__                 __                   .__  __  .__          ",
			"\\______   \\  | _____    ____ |  | __  ______ _____ |__|/  |_|  |__    /\\ ",
			" |    |  _/  | \\__  \\ _/ ___\\|  |/ / /  ___//     \\|  \\   __\\  |  \\   \\/ ",
			" |    |   \\  |__/ __ \\\\  \\___|    <  \\___ \\|  Y Y  \\  ||  | |   Y  \\  /\\ ",
			" |______  /____(____  /\\___  >__|_ \\/____  >__|_|  /__||__| |___|  /  \\/ ",
			"        \\/          \\/     \\/     \\/     \\/      \\/              \\/     \n\n" + "\033[0m",
		}
		DisplayMenu(blacksmith)
		Write("Weapon : \n\n	1 - GreatSword (x1 Iron Ore | x50 Zennys) \n	2 - LongSword (x1 Iron Ore | x50 Zennys) \n	3 - DualBlades (x1 Iron Ore | x50 Zennys) \n\n0 - Back\n\n")
		tobuy := ""
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			tobuy = "GreatSword"
		case "2":
			tobuy = "LongSword"
		case "3":
			tobuy = "DualBlades"
		default:
			Weapon(p)
		}
		if Ihave(p, player.Iron_ore, 1) && p.Money >= 100 {
			if tobuy == "GreatSword" {
				p.Weapon.Damage = 25
				p.Weapon.Speed = 5
				p.Weapon.Level = 1
			} else if tobuy == "LongSword" {
				p.Weapon.Damage = 15
				p.Weapon.Speed = 15
				p.Weapon.Level = 1
			} else {
				p.Weapon.Damage = 10
				p.Weapon.Speed = 25
				p.Weapon.Level = 1
			}
			p.Weapon.Name = "Iron"
			BuyUpgradeWeapon(p, &p.Weapon.Name, &tobuy, player.Iron_ore, 1)
		} else {
			Write("You don't got the materials buddy")
			time.Sleep(3 * time.Second)
		}
	case "2":
		if p.Weapon.Type == "Stick" || p.Weapon.Type == "Sexcalibur" {
			Write("You can't upgrade this weapon..")
			time.Sleep(3 * time.Second)
			Blacksmith(p)
		} else {
			Upgrade(p)
		}
	default:
		Blacksmith(p)
	}
}

func Upgrade(p *player.Character) {
	Clear()
	blacksmith := []string{
		"\033[95m" + "__________.__                 __                   .__  __  .__          ",
		"\\______   \\  | _____    ____ |  | __  ______ _____ |__|/  |_|  |__    /\\ ",
		" |    |  _/  | \\__  \\ _/ ___\\|  |/ / /  ___//     \\|  \\   __\\  |  \\   \\/ ",
		" |    |   \\  |__/ __ \\\\  \\___|    <  \\___ \\|  Y Y  \\  ||  | |   Y  \\  /\\ ",
		" |______  /____(____  /\\___  >__|_ \\/____  >__|_|  /__||__| |___|  /  \\/ ",
		"        \\/          \\/     \\/     \\/     \\/      \\/              \\/     \n\n" + "\033[0m",
	}
	DisplayMenu(blacksmith)
	if p.Weapon.Level == 1 && Ihave(p, player.Iron_ore, 1) && p.Money >= 100 {
		Write("Do you want to spend 1 Iron Ore and 100 zennys to upgrade your weapon ?\n\n	1 - Yes\n	2 - No")
		choice := ""
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			p.Weapon.Level++
			p.Weapon.Damage *= 2
			p.Weapon.Speed = int(float64(p.Weapon.Speed) * 1.5)
			name := "Machalite"
			BuyUpgrade(p, &p.Weapon.Name, &name, player.Iron_ore, 1)
		default:
			Weapon(p)
		}
	}
	if p.Weapon.Level == 2 && Ihave(p, player.Machalite_ore, 1) && p.Money >= 250 {
		Write("Do you want to spend 1 Machalite Ore and 250 zennys to upgrade your weapon ?\n\n	1 - Yes\n	2 - No")
		choice := ""
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			p.Weapon.Level++
			p.Weapon.Damage *= 2
			p.Weapon.Speed = int(float64(p.Weapon.Speed) * 1.5)
			name := "Dragonite"
			BuyUpgrade(p, &p.Weapon.Name, &name, player.Machalite_ore, 1)
		default:
			Weapon(p)
		}
	}
	if p.Weapon.Level == 3 && Ihave(p, player.Dragonite_ore, 1) && p.Money >= 500 {
		Write("Do you want to spend 1 Dragonite Ore and 500 zennys to upgrade your weapon ?\n\n	1 - Yes\n	2 - No")
		choice := ""
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			p.Weapon.Level++
			p.Weapon.Damage *= 2
			p.Weapon.Speed = int(float64(p.Weapon.Speed) * 1.5)
			name := "Carbalite"
			BuyUpgrade(p, &p.Weapon.Name, &name, player.Dragonite_ore, 1)
		default:
			Weapon(p)
		}
	}
	if p.Weapon.Level == 4 && Ihave(p, player.Carbalite_ore, 1) && p.Money >= 500 {
		Write("Do you want to spend 1 Carbalite Ore and 500 zennys to upgrade your weapon ?\n\n	1 - Yes\n	2 - No")
		choice := ""
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			p.Weapon.Level++
			p.Weapon.Damage *= 2
			p.Weapon.Speed = int(float64(p.Weapon.Speed) * 1.5)
			name := "Fucium"
			BuyUpgrade(p, &p.Weapon.Name, &name, player.Carbalite_ore, 1)
		default:
			Weapon(p)
		}
	}
	if p.Weapon.Level == 5 && Ihave(p, player.Fucium_ore, 1) && p.Money >= 1000 {
		Write("Do you want to spend 1 Fucium Ore and 1000 zennys to upgrade your weapon ?\n\n	1 - Yes\n	2 - No")
		choice := ""
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			p.Weapon.Level++
			p.Weapon.Damage *= 2
			p.Weapon.Speed = int(float64(p.Weapon.Speed) * 1.5)
			name := "Eltalite"
			BuyUpgrade(p, &p.Weapon.Name, &name, player.Fucium_ore, 1)
		default:
			Weapon(p)
		}
	}
	if p.Weapon.Level == 6 && Ihave(p, player.Eltalite_ore, 1) && p.Money >= 2500 {
		Write("Do you want to spend 1 Eltalite Ore and 2500 zennys to upgrade your weapon ?\n\n	1 - Yes\n	2 - No")
		choice := ""
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			p.Weapon.Level++
			p.Weapon.Damage *= 2
			p.Weapon.Speed = int(float64(p.Weapon.Speed) * 1.5)
			name := "Pure Dragon Blood"
			BuyUpgrade(p, &p.Weapon.Name, &name, player.Eltalite_ore, 1)
		default:
			Weapon(p)
		}
	} else {
		Write("Already maxed")
		time.Sleep(3 * time.Second)
		Blacksmith(p)
	}

}

func BuyUpgradeWeapon(p *player.Character, part *string, equipment *string, material player.Item, amount int) {
	Clear()
	blacksmith := []string{
		"\033[95m" + "__________.__                 __                   .__  __  .__          ",
		"\\______   \\  | _____    ____ |  | __  ______ _____ |__|/  |_|  |__    /\\ ",
		" |    |  _/  | \\__  \\ _/ ___\\|  |/ / /  ___//     \\|  \\   __\\  |  \\   \\/ ",
		" |    |   \\  |__/ __ \\\\  \\___|    <  \\___ \\|  Y Y  \\  ||  | |   Y  \\  /\\ ",
		" |______  /____(____  /\\___  >__|_ \\/____  >__|_|  /__||__| |___|  /  \\/ ",
		"        \\/          \\/     \\/     \\/     \\/      \\/              \\/     \n\n" + "\033[0m",
	}
	DisplayMenu(blacksmith)
	p.Inventory[material] -= amount
	p.Money -= material.Price * 2
	p.Weapon.Type = *equipment
	Write("Thanks brother")
	time.Sleep(time.Second)
	p.Max_health_point = player.HpUpdate(p.Armor, *p)
	player.SkillDetection(&p.Weapon)
	Blacksmith(p)
}

func Armor(p *player.Character) {
	Clear()
	blacksmith := []string{
		"\033[95m" + "__________.__                 __                   .__  __  .__          ",
		"\\______   \\  | _____    ____ |  | __  ______ _____ |__|/  |_|  |__    /\\ ",
		" |    |  _/  | \\__  \\ _/ ___\\|  |/ / /  ___//     \\|  \\   __\\  |  \\   \\/ ",
		" |    |   \\  |__/ __ \\\\  \\___|    <  \\___ \\|  Y Y  \\  ||  | |   Y  \\  /\\ ",
		" |______  /____(____  /\\___  >__|_ \\/____  >__|_|  /__||__| |___|  /  \\/ ",
		"        \\/          \\/     \\/     \\/     \\/      \\/              \\/     \n\n" + "\033[0m",
	}
	DisplayMenu(blacksmith)
	Write("Armors : \n\n")
	Write("	1 - Great Jagras (25hp per part)	\n	2 - Barroth (50hp per part)	\n	3 - Rathalos (100hp per part)	\n	4 - Nergigante (250hp per part)	\n	5 - Xeno'jiiva (500hp per part) \n\n	0 - Back\n\n")
	name := ""
	material := player.Iron_ore
	cost := 0
	choice := ""
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		name = "Great Jagras"
		material = player.Jagras_scale
		cost = 100
	case "2":
		name = "Barroth"
		material = player.Barroth_scalp
		cost = 200
	case "3":
		name = "Rathalos"
		material = player.Rathalos_cortex
		cost = 350
	case "4":
		name = "Nergigante"
		material = player.Nergigante_talon
		cost = 500
	case "5":
		name = "Xeno'jiiva"
		material = player.Xeno_jiiva_gem
		cost = 750
	case "0":
		Blacksmith(p)
	default:
		Write("not an answer..")
		Armor(p)
	}
	Clear()
	blacksmith = []string{
		"\033[95m" + "__________.__                 __                   .__  __  .__          ",
		"\\______   \\  | _____    ____ |  | __  ______ _____ |__|/  |_|  |__    /\\ ",
		" |    |  _/  | \\__  \\ _/ ___\\|  |/ / /  ___//     \\|  \\   __\\  |  \\   \\/ ",
		" |    |   \\  |__/ __ \\\\  \\___|    <  \\___ \\|  Y Y  \\  ||  | |   Y  \\  /\\ ",
		" |______  /____(____  /\\___  >__|_ \\/____  >__|_|  /__||__| |___|  /  \\/ ",
		"        \\/          \\/     \\/     \\/     \\/      \\/              \\/     \n\n" + "\033[0m",
	}
	DisplayMenu(blacksmith)
	armors := []string{
		"1 - Helm " + name + " : " + material.Name + " x1 & zenny x" + strconv.Itoa(cost),
		"2 - Chest " + name + " : " + material.Name + " x1 & zenny x" + strconv.Itoa(cost),
		"3 - Arms " + name + " : " + material.Name + " x1 & zenny x" + strconv.Itoa(cost),
		"4 - Waist " + name + " : " + material.Name + " x1 & zenny x" + strconv.Itoa(cost),
		"5 - Legs " + name + " : " + material.Name + " x1 & zenny x" + strconv.Itoa(cost) + "\n",
		"6 - Back",
		"0 - Leave\n"}
	DisplayMenu(armors)
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		BuyUpgrade(p, &p.Armor.Helm, &name, material, 1)
	case "2":
		BuyUpgrade(p, &p.Armor.Chest, &name, material, 1)
	case "3":
		BuyUpgrade(p, &p.Armor.Arms, &name, material, 1)
	case "4":
		BuyUpgrade(p, &p.Armor.Waist, &name, material, 1)
	case "5":
		BuyUpgrade(p, &p.Armor.Legs, &name, material, 1)
	case "0":
		MainMenu(p)
	case "6":
		Armor(p)
	}
}

func BuyUpgrade(p *player.Character, part *string, equipment *string, material player.Item, amount int) {
	Clear()
	blacksmith := []string{
		"\033[95m" + "__________.__                 __                   .__  __  .__          ",
		"\\______   \\  | _____    ____ |  | __  ______ _____ |__|/  |_|  |__    /\\ ",
		" |    |  _/  | \\__  \\ _/ ___\\|  |/ / /  ___//     \\|  \\   __\\  |  \\   \\/ ",
		" |    |   \\  |__/ __ \\\\  \\___|    <  \\___ \\|  Y Y  \\  ||  | |   Y  \\  /\\ ",
		" |______  /____(____  /\\___  >__|_ \\/____  >__|_|  /__||__| |___|  /  \\/ ",
		"        \\/          \\/     \\/     \\/     \\/      \\/              \\/     \n\n" + "\033[0m",
	}
	DisplayMenu(blacksmith)
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
			p.Inventory[material] -= amount
			p.Money -= material.Price * 2
			*part = *equipment
			Write("Thanks brother")
			time.Sleep(time.Second)
			p.Max_health_point = player.HpUpdate(p.Armor, *p)
			player.SkillDetection(&p.Weapon)
			Blacksmith(p)
		}
	}
}

func Ihave(p *player.Character, material player.Item, amount int) bool {
	return p.Inventory[material] >= amount
}
