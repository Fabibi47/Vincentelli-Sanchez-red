package menu

import (
	"fmt"
	"player"
	"strconv"
)

func MenuInventory(p player.Character) {
	inventorySlice := []string{"INVENTORY \n\n\n"}
	for object, quantity := range p.Inventory {
		inventorySlice = append(inventorySlice, object+strconv.Itoa(quantity))
	}
	DisplayMenu(inventorySlice)
	action := ""
	fmt.Scanln(&action)
	switch action {
	default:
		MainMenu(p)
	}
}
