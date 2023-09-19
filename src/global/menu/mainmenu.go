package menu

import "os"

func MainMenu() {
	var mainMenu []string = []string{"1 - Character \n", "2 - Inventory \n", "3 - Shop \n", "4 - Leave"}
	action := DisplayMenu(mainMenu)
	switch action {
	case "1":
	case "2":
	case "3":
		Marchand()
	case "4":
		Write("Are you sure ? \n \n 1 - Yes, leave\n 2 - No, continue\n")
		scanner.Scan()
		response := scanner.Text()
		switch response {
		case "1":
			os.Stdout.WriteString("Left")
		case "2":
			DisplayMenu(mainMenu)
		}
	}
}
