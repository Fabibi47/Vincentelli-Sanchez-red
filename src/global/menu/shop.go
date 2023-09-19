package menu

func Marchand() {
	var marchand []string = []string{"Welcome to my shop! Feel free to take a look! \n \n", "0 - Nevermind, bye. (Leave)\n", "1 - Potion: 50 Zennys", "2 - Mega Potion: 150 Zennys"}
	a := map[string]int{
		"1 - Potion":      50,
		"2 - Mega Potion": 150,
	}
	for article := range a {
		marchand = append(marchand, article)
	}
	action := DisplayMenu(marchand)
	if action == "0" {
		MainMenu()
	} else if action == "1" {
		Write("Thank you for ")
	} else if action == "2" {

	} else {
		Write("Sorry, I don't have this. Choose a correct number.")
		Marchand()
	}
}
