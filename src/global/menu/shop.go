package menu

import (
	"fmt"
	"player"
)

func Marchand(p *player.Character) {
	var marchand []string = []string{
		"Welcome to my shop! How can I help you? \n \n",
		"0 - Nevermind, bye. (Leave)\n",
		"1 - I want to buy something",
		"2 - I want to sell something"}
	Clear()
	DisplayMenu(marchand)
	navigating := true
	scanner.Scan()
	for navigating {
		answer := scanner.Text()
		switch answer {
		case "0":
			navigating = false
			MainMenu(p)
		case "1":
			Buy(p)
		case "2":
			Sell(p)
		}
	}
}

func Buy(p *player.Character) {
	buying := []string{"Feel free to take a look !\n",
		"0 - I changed my mind (back)\n",
		"1 - Potion: 50 Zennys",
		"2 - Mega Potion: 150 Zennys",
		"3 - Poison: 100 Zennys"}
	Clear()
	shopping := true
	DisplayMenu(buying)
	for shopping {
		scanner.Scan()
		action := scanner.Text()
		switch action {
		case "0":
			shopping = false
			Marchand(p)
		case "1":
			if p.Money >= player.Potion.Price {
				p.Inventory[player.Potion]++
				p.Money -= player.Potion.Price
				Write("Thank you, Hunter ! \nAnything else?")
			} else {
				Write("Sorry, Hunter, but you don't have enough Zennys for this item.")
			}
		case "2":
			if p.Money >= player.Mega_Potion.Price {
				p.Inventory[player.Mega_Potion]++
				p.Money -= player.Mega_Potion.Price
				Write("Thank you, Hunter ! \nAnything else?")
			} else {
				Write("Sorry, Hunter, but you don't have enough Zennys for this item.")
			}
		case "3":
			if p.Money >= player.Poison.Price {
				p.Inventory[player.Poison]++
				p.Money -= player.Poison.Price
				Write("Thank you, Hunter ! \nAnything else?")
			} else {
				Write("Sorry, Hunter, but you don't have enough Zennys for this item.")
			}
		default:
			Write("Sorry, I don't have this kind of thing.")
		}
	}
}

func Sell(p *player.Character) {
	sell := []string{"What do you want to sell?",
		"0 - I changed my mind (back)"}
	articles := []player.Item{}
	for a := range p.Inventory {
		articles = append(articles, a)
	}
	for i := range articles {
		sell = append(sell, fmt.Sprint(i+1)+"-"+articles[i].Name)
	}
	selling := true
	Clear()
	DisplayMenu(sell)
	for selling {
		scanner.Scan()
		action := scanner.Text()
		switch action {
		case "1":

		}
	}
}
