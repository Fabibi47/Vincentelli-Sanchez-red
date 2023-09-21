package menu

import (
	"os"
	"player"
)

func MainMenu(p player.Character) {
	var mainMenu []string = []string{"1 - Character \n", "2 - Inventory \n", "3 - Shop \n", "4 - Blacksmith \n", "5 - Battle \n", "6 - Quests \n", "7 - Leave \n\n"}
	DisplayMenu(mainMenu)
	navigating := true
	scanner.Scan()
	for navigating {
		action := scanner.Text()
		switch action {
		case "1":
			navigating = false
			Stats(p)
		case "2":
			navigating = false
			MenuInventory(p)
		case "3":
			navigating = false
			Marchand(p)
		case "4":
			navigating = false
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
			navigating = true
		}
	}
}
