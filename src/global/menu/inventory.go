package menu

import (
	"fmt"
	"red/global/player"
	"strconv"
)

func MenuInventory(p *player.Character) {
	inventorySlice := []string{"INVENTORY \n\n\n"}
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
