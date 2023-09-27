package menu

import (
	"fmt"
	"player"
)

func MenuInventory(p player.Character) {
	inventorySlice := []string{"INVENTORY \n\n\n"}
	for object, _ := range p.Inventory {
		inventorySlice = append(inventorySlice, object.Name)
	}
	DisplayMenu(inventorySlice)
	action := ""
	fmt.Scanln(&action)
	switch action {
	default:
		MainMenu(p)
	}
}
