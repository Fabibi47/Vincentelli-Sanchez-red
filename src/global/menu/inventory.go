package menu

import (
	"fmt"
	"red/global/player"
	"strconv"
)

func MenuInventory(p *player.Character) {
	inventorySlice := []string{
		"" + "\033[33m" + " _  " + "\033[0m" + "                     _                   _ ",
		"" + "\033[33m" + "| |._ _  _ _" + "\033[0m" + "  ___ ._ _ _| |_ ___  _ _  _ _  <_>",
		"" + "\033[33m" + "| || ' || | |/ ._>| ' " + "\033[0m" + "| | | / . \\| '_>| | |  _ ",
		"" + "\033[33m" + "|_||_|_||__/ \\___.|_|_| |_| \\" + "\033[0m" + "___/|_|  `_. | <_>",
		"" + "\033[33m" + "                                      <___'  " + "\033[0m" + "  ",
		"\n\n\n"}
	for object := range p.Inventory {
		inventorySlice = append(inventorySlice, "	"+object.Name+" x"+strconv.Itoa(p.Inventory[object]))
	}
	Clear()
	DisplayMenu(inventorySlice)
	action := ""
	fmt.Scanln(&action)
	switch action {
	default:
		MainMenu(p)
	}
}
