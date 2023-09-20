package menu

import (
	"fmt"
	"player"
)

func Marchand(p player.Character) {
	var marchand []string = []string{
		"Welcome to my shop! How can I help you? \n \n",
		"0 - Nevermind, bye. (Leave)\n",
		"1 - I want to buy something",
		"2 - I want to sell something"}
	DisplayMenu(marchand)
	navigating := true
	scanner.Scan()
	for navigating {
		answer := scanner.Text()
		switch answer {
		case "0":
			MainMenu(p)
		case "1":
			Buy(p)
		case "2":
			Sell(p)
		}
	}
}

func Buy(p player.Character) {
	buying := []string{"Feel free to take a look !\n",
		"0 - I changed my mind (back)\n",
		"1 - Potion: 50 Zennys",
		"2 - Mega Potion: 150 Zennys",
		"3 - Poison: 100 Zennys"}
	a := map[string]int{
		"Potion":      50,
		"Mega Potion": 150,
		"Poison":      100,
	}
	shopping := true
	for article := range a {
		buying = append(buying, article)
	}
	DisplayMenu(buying)
	scanner.Scan()
	for shopping {
		action := scanner.Text()
		switch action {
		case "0":
			shopping = false
			Marchand(p)
		case "1":
			if p.Money >= a["Potion"] {
				p.Inventory["Potion"]++
				p.Money -= a["Potion"]
				Write("Thank you, Hunter ! \nAnything else?")
			} else {
				Write("Sorry, Hunter, but you don't have enough Zennys for this item.")
			}
		case "2":
			if p.Money >= a["Mega Potion"] {
				p.Inventory["Mega Potion"]++
				p.Money -= a["Mega Potion"]
				Write("Thank you, Hunter ! \nAnything else?")
			} else {
				Write("Sorry, Hunter, but you don't have enough Zennys for this item.")
			}
		case "3":
			if p.Money >= a["Poison"] {
				p.Inventory["Poison"]++
				p.Money -= a["Poison"]
				Write("Thank you, Hunter ! \nAnything else?")
			} else {
				Write("Sorry, Hunter, but you don't have enough Zennys for this item.")
			}
		default:
			Write("Sorry, I don't have this kind of thing.")
		}
	}
}

func Sell(p player.Character) {
	sell := []string{"What do you want to sell?",
		"0 - I changed my mind (back)"}
	articles := []string{}
	for a := range p.Inventory {
		articles = append(articles, a)
	}
	for i := range articles {
		sell = append(sell, fmt.Sprint(i+1)+"-"+articles[i])
	}
	selling := true
	DisplayMenu(sell)
	scanner.Scan()
	for selling {
		action := scanner.Text()
		switch action {
		case "1":
		}
	}
}
