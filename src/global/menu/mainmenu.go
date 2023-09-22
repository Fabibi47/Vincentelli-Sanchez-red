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
	navigating := true
	scanner.Scan()
	for navigating {
		action := scanner.Text()
		switch action {
		case "3":
			navigating = false
			Stats(p)
		case "4":
			navigating = false
			MenuInventory(p)
		case "5":
			navigating = false
			Marchand(p)
		case "6":
			navigating = false
			Blacksmith(p)
		case "7":
			Write("Are you sure ? \n \n 1 - Yes, leave\n 2 - No, continue\n")
			scanner.Scan()
			response := scanner.Text()
			switch response {
			case "1":
				os.Stdout.WriteString("Left")
				navigating = false
			case "2":
				MainMenu(p)
			}
		default:
			navigating = true
		}
	}
}
