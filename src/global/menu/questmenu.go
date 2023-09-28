package menu

import (
	"fmt"
	"monsters"
	"player"
)

func QuestMenu(p player.Character) {
	Write("QUESTS")
	questSlide := []string{
		"1 - Kill a Great Jagras\n",
		"2 - Kill a Barroth\n",
		"3 - Kill a Rathalos\n",
		"4 - Kill a Nergigante\n",
		"5 - Kill a Xeno'Jiiva\n",
		"\n0 - Back",
	}
	DisplayMenu(questSlide)
	action := ""
	fmt.Scanln(&action)
	switch action {
	case "1":
		BossBattle(&p, &monsters.Great_Jagras)
	case "2":
		BossBattle(&p, &monsters.Barroth)
	case "3":
		BossBattle(&p, &monsters.Rathalos)
	case "4":
		BossBattle(&p, &monsters.Nergigante)
	case "5":
		BossBattle(&p, &monsters.Xeno_Jiiva)
	default:
		MainMenu(&p)
	}
}
