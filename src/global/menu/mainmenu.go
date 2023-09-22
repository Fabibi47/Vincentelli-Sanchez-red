package menu

import (
	"os"
	"player"
)

func MainMenu(p player.Character) {
	var mainMenu []string = []string{
		"1 - Battle \n",
		"2 - Quests \n",
		"3 - Character \n",
		"4 - Inventory \n",
		"5 - Shop \n",
		"6 - Blacksmith \n",
		"7 - Leave \n\n"}
	Clear()
	DisplayMenu(mainMenu)
	scanner.Scan()
	action := scanner.Text()
	switch action {
	case "3":
		Stats(p)
	case "4":
		MenuInventory(p)
	case "5":
		Marchand(p)
	case "6":
		Blacksmith(p)
	case "7":
		Write("Are you sure ? \n \n 1 - Yes, leave\n 2 - No, continue\n")
		scanner.Scan()
		response := scanner.Text()
		switch response {
		case "1":
			os.Stdout.WriteString("Left")
		case "2":
			MainMenu(p)
		}
	default:
		MainMenu(p)
	}
}
