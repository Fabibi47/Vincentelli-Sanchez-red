package menu

import (
	"fmt"
	"player"
	"strconv"
	"time"
)

func MenuInventory(p player.Character) {
	Clear()
	Write("INVENTORY \n\n\n")
	inventorySlice := []string{}
	items := []player.Item{}
	for item, quantity := range p.Inventory {
		inventorySlice = append(inventorySlice, item.Name+" x"+strconv.Itoa(quantity))
		items = append(items, item)
	}
	DisplayMenu(inventorySlice)
	Write("1 - Navigate")
	Write("0 - Back")
	action := ""
	fmt.Scanln(&action)
	switch action {
	case "1":
		selectionSlice := []string{}
		for j, i := range items {
			selectionSlice = append(selectionSlice, strconv.Itoa(j+1)+" - "+i.Name)
		}
		DisplayMenu(selectionSlice)
		Write("0 - Back")
		fmt.Scanln(&action)
		if action > strconv.Itoa(0) && action <= strconv.Itoa(len(selectionSlice)) {
			item, _ := strconv.Atoi(action)
			item--
			DisplayObject(items[item], p)
		}
	default:
		MainMenu(p)
	}
}

func DisplayObject(item player.Item, p player.Character) {
	Clear()
	action := ""
	Write(item.Name + "\n\n")
	Write(item.Descritpion)
	if item.Usable {
		Write("1 - Use")
	}
	Write("0 - Back")
	fmt.Scanln(&action)
	switch action {
	case "1":
		if item.Usable {
			item.Use(&p)
			Write(item.Name + " used succesfully !")
			time.Sleep(3 * time.Second)
			MenuInventory(p)
		} else {
			MenuInventory(p)
		}
	default:
		MenuInventory(p)
	}
}
