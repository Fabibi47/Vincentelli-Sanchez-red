package menu

import (
	"fmt"
	"monsters"
	"player"
	"strconv"
	"time"
)

func Victory(p *player.Character, m monsters.Monster, zone string) {
	Clear()
	frame1 := []string{"\033[33m" +
		"┓┏       ┏┓       ┓      ┓",
		"┣┫┓┏┏┓╋  ┃ ┏┓┏┳┓┏┓┃┏┓╋┏┓┏┫",
		"┛┗┗┛┛┗┗  ┗┛┗┛┛┗┗┣┛┗┗ ┗┗ ┗┻",
		"\033[0m",
	}
	DisplayMenu(frame1)
	Write("\n\n\nYou earned " + strconv.Itoa(m.Drop) + " zennys and " + strconv.Itoa(m.Exp) + " experience!\n\n")
	p.Money += m.Drop
	p.Exp += m.Exp
	player.LevelUpdate(p)
	time.Sleep(1 * time.Second)
	Write("1 - Next monster\n")
	Write("0 - Back\n")
	action := ""
	fmt.Scanln(&action)
	switch action {
	case "1":
		Hunt(p, GetMonster(zone), zone)
	default:
		MainMenu(p)
	}
}
