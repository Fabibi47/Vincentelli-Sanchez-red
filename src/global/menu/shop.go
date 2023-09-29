package menu

import (
	"red/global/player"
	"time"
)

func Marchand(p *player.Character) {
	var marchand []string = []string{
		"   .-'''-. .---.  .---.     ,-----.    .-------.         " + "\033[31m" + "_ _" + "\033[0m" + "   ",
		"  / _     \\|   |  |_ _|   .'  .-,  '.  \\  " + "\033[96m" + "_(`)_" + "\033[0m" + " \\       " + "\033[31m" + "( ` )" + "\033[0m" + "  ",
		" " + "\033[91m" + "(`' )" + "\033[0m" + "/`--'|   |  " + "\033[94m" + "( ' )" + "\033[0m" + "  / ,-.|  \\ _ \\ | " + "\033[96m" + "(_ " + "\033[33m" + "o" + "\033[96m" + "._)" + "\033[0m" + "|      " + "\033[31m" + "(_" + "\033[33m" + "{;}" + "\033[31m" + "_)" + "\033[0m" + " ",
		"" + "\033[91m" + "(_ " + "\033[33m" + "o" + "\033[91m" + " _)" + "\033[0m" + ".   |   '-" + "\033[94m" + "(_" + "\033[33m" + "{;}" + "\033[94m" + "_)" + "\033[0m" + ";  \\  '_ /  | :|  " + "\033[96m" + "(_,_)" + "\033[0m" + " /       " + "\033[31m" + "(_,_)" + "\033[0m" + "  ",
		" " + "\033[91m" + "(_,_)" + "\033[0m" + ". '. |      " + "\033[94m" + "(_,_)" + "\033[0m" + " |  _`,/ \\ _/  ||   '-.-'               ",
		".---.  \\  :| _ _--.   | : (  '\\_/ \\   ;|   |              " + "\033[34m" + "_" + "\033[0m" + "    ",
		"\\    `-'  ||" + "\033[92m" + "( ' )" + "\033[0m" + " |   |  \\ `\"/  \\  ) / |   |            " + "\033[34m" + "_( )_" + "\033[0m" + "  ",
		" \\       / " + "\033[92m" + "(_" + "\033[33m" + "{;}" + "\033[92m" + "_)" + "\033[0m" + "|   |   '. \\_/``\".'  /   )           " + "\033[34m" + "(_ " + "\033[33m" + "o" + "\033[34m" + " _)" + "\033[0m" + " ",
		"  `-...-'  '" + "\033[92m" + "(_,_)" + "\033[0m" + " '---'     '-----'    `---'            " + "\033[34m" + "(_,_)" + "\033[0m" + "  \n\n\n",
		"Welcome to my shop! How can I help you? \n\n",
		"1 - I want to buy something",
		"0 - Nevermind, bye. (Leave)\n",
	}
	Clear()
	DisplayMenu(marchand)
	navigating := true
	scanner.Scan()
	for navigating {
		answer := scanner.Text()
		switch answer {
		default:
			navigating = false
			MainMenu(p)
		case "1":
			Buy(p)
		}
	}
}

func Buy(p *player.Character) {
	buying := []string{"Feel free to take a look !\n",
		"0 - I changed my mind (back)\n",
		"1 - Potion: 50 Zennys",
		"2 - Mega Potion: 150 Zennys",
		"3 - Poison: 100 Zennys",
		"4 - Antidote: 150 Zennys",
		"5 - Backpack: 250 Zennys",
		"6 - Iron Ore : 50 Zennys",
		"7 - Malachite Ore : 100 Zennys",
		"8 - Dragonite Ore : 200 Zennys",
		"9 - Carbalite Ore : 500 Zennys",
		"10 - Fucium Ore : 1000 Zennys",
		"11 - Eltalite Ore : 2000 Zennys",
	}
	Clear()
	shopping := true
	for shopping {
		Clear()
		DisplayMenu(buying)
		scanner.Scan()
		action := scanner.Text()
		switch action {
		case "0":
			shopping = false
			Marchand(p)
		case "1":
			if p.Money >= player.Potion.Price && p.Inventory[player.Potion] < p.Stack {
				p.Inventory[player.Potion]++
				p.Money -= player.Potion.Price
				Write("Thank you," + "\033[31m" + " Hunter" + "\033[0m" + " ! \nAnything else?")
			} else if p.Inventory[player.Potion] >= p.Stack {
				Write("Your bag is full, you can't take anymore potions.")
			} else {
				Write("Sorry," + "\033[31m" + " Hunter" + "\033[0m" + ", but you don't have enough Zennys for this item.")
			}
		case "2":
			if p.Money >= player.Mega_Potion.Price && p.Inventory[player.Mega_Potion] < p.Stack {
				p.Inventory[player.Mega_Potion]++
				p.Money -= player.Mega_Potion.Price
				Write("Thank you," + "\033[31m" + " Hunter" + "\033[0m" + " ! \nAnything else?")
			} else if p.Inventory[player.Mega_Potion] >= p.Stack {
				Write("Your bag is full, you can't take anymore nega potions.")
			} else {
				Write("Sorry," + "\033[31m" + " Hunter" + "\033[0m" + ", but you don't have enough Zennys for this item.")
			}
		case "3":
			if p.Money >= player.Poison.Price && p.Inventory[player.Poison] < p.Stack {
				p.Inventory[player.Poison]++
				p.Money -= player.Poison.Price
				Write("Thank you, Hunter ! \nAnything else?")
			} else if p.Inventory[player.Poison] >= p.Stack {
				Write("Your bag is full, you can't take anymore poisons.")
			} else {
				Write("Sorry," + "\033[31m" + " Hunter" + "\033[0m" + ", but you don't have enough Zennys for this item.")
			}
		case "4":
			if p.Money >= player.Antidote.Price && p.Inventory[player.Antidote] < p.Stack {
				p.Inventory[player.Antidote]++
				p.Money -= player.Antidote.Price
				Write("Thank you, Hunter ! \nAnything else?")
			} else if p.Inventory[player.Antidote] >= p.Stack {
				Write("Your bag is full, you can't take anymore poisons.")
			} else {
				Write("Sorry," + "\033[31m" + " Hunter" + "\033[0m" + ", but you don't have enough Zennys for this item.")
			}
		case "5":
			if p.Stack >= 10 {
				Write("You already have this item.")
			} else if p.Money >= 250 {
				Write("Thank you," + "\033[31m" + " Hunter" + "\033[0m" + " ! \nAnything else?")
				p.Stack = 10
			} else {
				Write("Sorry," + "\033[31m" + " Hunter" + "\033[0m" + ", but you don't have enough Zennys for this item.")
			}
		case "6":
			if p.Money >= player.Iron_ore.Price {
				p.Inventory[player.Iron_ore] += 1
				p.Money -= player.Iron_ore.Price
				Write("Thank you, Hunter ! \nAnything else?")
			} else {
				Write("Sorry," + "\033[31m" + " Hunter" + "\033[0m" + ", but you don't have enough Zennys for this item.")
			}
		case "7":
			if p.Money >= player.Machalite_ore.Price {
				p.Inventory[player.Machalite_ore] += 1
				p.Money -= player.Machalite_ore.Price
				Write("Thank you, Hunter ! \nAnything else?")
			} else {
				Write("Sorry," + "\033[31m" + " Hunter" + "\033[0m" + ", but you don't have enough Zennys for this item.")
			}
		case "8":
			if p.Money >= player.Dragonite_ore.Price {
				p.Inventory[player.Dragonite_ore] += 1
				p.Money -= player.Dragonite_ore.Price
				Write("Thank you, Hunter ! \nAnything else?")
			} else {
				Write("Sorry," + "\033[31m" + " Hunter" + "\033[0m" + ", but you don't have enough Zennys for this item.")
			}
		case "9":
			if p.Money >= player.Carbalite_ore.Price {
				p.Inventory[player.Carbalite_ore] += 1
				p.Money -= player.Carbalite_ore.Price
				Write("Thank you, Hunter ! \nAnything else?")
			} else {
				Write("Sorry," + "\033[31m" + " Hunter" + "\033[0m" + ", but you don't have enough Zennys for this item.")
			}
		case "10":
			if p.Money >= player.Fucium_ore.Price {
				p.Inventory[player.Fucium_ore] += 1
				p.Money -= player.Fucium_ore.Price
				Write("Thank you, Hunter ! \nAnything else?")
			} else {
				Write("Sorry," + "\033[31m" + " Hunter" + "\033[0m" + ", but you don't have enough Zennys for this item.")
			}
		case "11":
			if p.Money >= player.Eltalite_ore.Price {
				p.Inventory[player.Eltalite_ore] += 1
				p.Money -= player.Eltalite_ore.Price
				Write("Thank you, Hunter ! \nAnything else?")
			} else {
				Write("Sorry," + "\033[31m" + " Hunter" + "\033[0m" + ", but you don't have enough Zennys for this item.")
			}
		default:
			Write("Sorry, I don't have this kind of thing.")
		}
		time.Sleep(time.Second)
	}
}
