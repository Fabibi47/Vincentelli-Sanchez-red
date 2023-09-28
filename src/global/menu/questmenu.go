package menu

import (
	"fmt"
	"monsters"
	"player"
	"time"
)

func QuestMenu(p player.Character) {
	Write("QUESTS")
	questSlide := []string{
		"Hey ! Good to see you, " + "\033[31m" + "Hunter" + "\033[0m" + ". Which quest do you want to do?",
		"1 - Kill a Great Jagras\n",
		"2 - Kill a Barroth\n",
		"3 - Kill a Rathalos\n",
		"4 - Kill a Nergigante\n",
		"5 - Kill a Xeno'Jiiva\n",
		"\n0 - Back",
	}
	if p.Quest == 5 {
		questSlide = append(questSlide, "By the way, we just received a new Quest!")
		questSlide = append(questSlide, "Seems really hard but I feel like you can do it.")
	}
	Clear()
	DisplayMenu(questSlide)
	action := ""
	fmt.Scanln(&action)
	switch action {
	case "1":
		Write1("I believe in you, " + "\033[31m" + "Hunter" + "\033[0m" + "!")
		time.Sleep(time.Second)
		BossBattle(&p, &monsters.Great_Jagras)
	case "2":
		if p.Quest == 1 {
			Write1("I believe in you, " + "\033[31m" + "Hunter" + "\033[0m" + "!")
			time.Sleep(time.Second)
			BossBattle(&p, &monsters.Barroth)
		} else {
			Write1("I don't want to be rude, but... You should try an easier one before.")
			time.Sleep(time.Second)
			QuestMenu(p)
		}
	case "3":
		if p.Quest == 2 {
			Write1("I believe in you, " + "\033[31m" + "Hunter" + "\033[0m" + "!")
			time.Sleep(time.Second)
			BossBattle(&p, &monsters.Rathalos)
		} else {
			Write1("I don't want to be rude, but... You should try an easier one before.")
			time.Sleep(time.Second)
			QuestMenu(p)
		}
	case "4":
		if p.Quest == 3 {
			Write1("I believe in you, " + "\033[31m" + "Hunter" + "\033[0m" + "!")
			time.Sleep(time.Second)
			BossBattle(&p, &monsters.Nergigante)
		} else {
			Write1("I don't want to be rude, but... You should try an easier one before.")
			time.Sleep(time.Second)
			QuestMenu(p)
		}
	case "5":
		if p.Quest == 4 {
			Write1("I believe in you, " + "\033[31m" + "Hunter" + "\033[0m" + "!")
			time.Sleep(time.Second)
			BossBattle(&p, &monsters.Xeno_Jiiva)
		} else {
			Write1("I don't want to be rude, but... You should try an easier one before.")
			time.Sleep(time.Second)
			QuestMenu(p)
		}
	case "6":
		if p.Quest == 5 {
			Write1("I believe in you, " + "\033[31m" + "Hunter" + "\033[0m" + "!")
			time.Sleep(time.Second)
			BossBattle(&p, &monsters.Alatreon)
		} else {
			Write1("See you later !")
			time.Sleep(time.Second)
			MainMenu(&p)
		}
	default:
		Write1("See you later !")
		time.Sleep(time.Second)
		MainMenu(&p)
	}
}
