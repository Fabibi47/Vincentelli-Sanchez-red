package menu

import (
	"fmt"
	"red/global/player"
	"strconv"
	"time"
)

func MenuInventory(p *player.Character) {
	inventorySlice := []string{
		"" + "\033[33m" + " _  " + "\033[0m" + "                     _                   _ ",
		"" + "\033[33m" + "| |._ _  _ _" + "\033[0m" + "  ___ ._ _ _| |_ ___  _ _  _ _  <_>",
		"" + "\033[33m" + "| || ' || | |/ ._>| ' " + "\033[0m" + "| | | / . \\| '_>| | |  _ ",
		"" + "\033[33m" + "|_||_|_||__/ \\___.|_|_| |_| \\" + "\033[0m" + "___/|_|  `_. | <_>",
		"" + "\033[33m" + "                                      <___'  " + "\033[0m" + "  ",
		"\n\n"}
	for object := range p.Inventory {
		inventorySlice = append(inventorySlice, "	"+object.Name+" x"+strconv.Itoa(p.Inventory[object]))
	}
	Clear()
	DisplayMenu(inventorySlice)
	Write("\n\n1 - Navigate")
	Write("\n0 - Back")
	action := ""
	fmt.Scanln(&action)
	switch action {
	default:
		MainMenu(p)
	case "1":
		Navigate(p)
	}
}

func Navigate(p *player.Character) {
	inventory := []player.Item{}
	for object := range p.Inventory {
		inventory = append(inventory, object)
	}
	navigateMenu := []string{"" + "\033[33m" + " _  " + "\033[0m" + "                     _                   _ ",
		"" + "\033[33m" + "| |._ _  _ _" + "\033[0m" + "  ___ ._ _ _| |_ ___  _ _  _ _  <_>",
		"" + "\033[33m" + "| || ' || | |/ ._>| ' " + "\033[0m" + "| | | / . \\| '_>| | |  _ ",
		"" + "\033[33m" + "|_||_|_||__/ \\___.|_|_| |_| \\" + "\033[0m" + "___/|_|  `_. | <_>",
		"" + "\033[33m" + "                                      <___'  " + "\033[0m" + "  ",
		"\n\n",
		"Navigating \n\n",
	}
	for i, o := range inventory {
		navigateMenu = append(navigateMenu, "	"+strconv.Itoa(i+1)+" - "+o.Name)
	}
	Clear()
	DisplayMenu(navigateMenu)
	Write("\n\n0 - Back")
	action := ""
	fmt.Scanln(&action)
	response, _ := strconv.Atoi(action)
	if response <= len(inventory) && action > "0" {
		DisplayObject(inventory[response-1], p)
	} else {
		MenuInventory(p)
	}
}

func DisplayObject(o player.Item, p *player.Character) {
	display := []string{"" + "\033[33m" + " _  " + "\033[0m" + "                     _                   _ ",
		"" + "\033[33m" + "| |._ _  _ _" + "\033[0m" + "  ___ ._ _ _| |_ ___  _ _  _ _  <_>",
		"" + "\033[33m" + "| || ' || | |/ ._>| ' " + "\033[0m" + "| | | / . \\| '_>| | |  _ ",
		"" + "\033[33m" + "|_||_|_||__/ \\___.|_|_| |_| \\" + "\033[0m" + "___/|_|  `_. | <_>",
		"" + "\033[33m" + "                                      <___'  " + "\033[0m" + "  ",
		"\n\n\n"}
	display = append(display, o.Name+"\n\n"+"   "+o.Descritpion+"\n")
	Clear()
	DisplayMenu(display)
	if o.Usable {
		Write("1 - Use")
	}
	Write("0 - Back")
	action := ""
	fmt.Scanln(&action)
	switch action {
	case "1":
		if o.Usable {
			o.Use(p)
			Write("Object used succesfully !")
			time.Sleep(time.Second)
		}
		Navigate(p)
	default:
		Navigate(p)
	}
}
