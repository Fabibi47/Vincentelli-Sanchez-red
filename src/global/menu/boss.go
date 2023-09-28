package menu

import (
	"fmt"
	"math/rand"
	"monsters"
	"player"
	"strconv"
	"time"
)

func BossBattle(p *player.Character, b *monsters.Boss) {
	playerHP := &p.Health_point
	bossHP := b.PV
	hunting := true
	playerAction := []string{
		"\n\n	1 - Attack",
		"	2 - Objects",
		"	3 - Run\n\n"}
	speedp := 0
	speedm := 0
	for hunting {
		speedp += p.Weapon.Speed
		speedm += b.Speed
		if speedp >= 1000 {
			speedp -= 1000
			battleMenu := []string{
				"Battle : \n\n\n",
				"   " + b.Name + "   " + strconv.Itoa(bossHP) + "/" + strconv.Itoa(b.PV) + "\n\n",
				" ↳ " + p.Name + "   " + strconv.Itoa(*playerHP) + "/" + strconv.Itoa(p.Max_health_point)}
			Clear()
			DisplayMenu(battleMenu)
			DisplayMenu(playerAction)
			action := ""
			fmt.Scanln(&action)
			switch action {
			case "1":
				Clear()
				DisplayMenu(battleMenu)
				DisplayMenu(GetSkills(p))
				attack := ""
				fmt.Scanln(&attack)
				switch attack {
				case "1":
					bossHP -= p.Weapon.Skills[0].Use(p)
					Write("Skill used !")
					time.Sleep(time.Second)
					if bossHP <= 0 {
						BossVictory(*p, *b)
						hunting = false
					}
				case "2":
					bossHP -= p.Weapon.Skills[1].Use(p)
					Write("Skill used !")
					time.Sleep(time.Second)
					if bossHP <= 0 {
						BossVictory(*p, *b)
						hunting = false
					}
				case "3":
					bossHP -= p.Weapon.Skills[2].Use(p)
					Write("Skill used !")
					time.Sleep(time.Second)
					if bossHP <= 0 {
						BossVictory(*p, *b)
						hunting = false
					}
				}
			case "2":
				Clear()
				DisplayMenu(battleMenu)
				Items(p)
			case "3":
				Clear()
				Write("You ran away... Loser.")
				time.Sleep(3 * time.Second)
				hunting = false
			}
		} else if speedm >= 1000 {
			speedm -= 1000
			time.Sleep(1 * time.Second)
			battleMenu := []string{
				"Battle : \n\n\n",
				" ↱ " + b.Name + "   " + strconv.Itoa(bossHP) + "/" + strconv.Itoa(b.PV) + "\n\n",
				"   " + p.Name + "   " + strconv.Itoa(*playerHP) + "/" + strconv.Itoa(p.Max_health_point)}
			Clear()
			DisplayMenu(battleMenu)
			fmt.Println("Boss's turn !")
			chance := rand.Int()
			if chance%5 == 0 {
				*playerHP -= b.Attacks[1].Use(*b)
				Write(b.Name + " used " + b.Attacks[1].Name)
			} else {
				*playerHP -= b.Attacks[0].Use(*b)
				Write(b.Name + " attacked !")
			}
			time.Sleep(time.Second)
			if *playerHP <= 0 {
				hunting = false
				Wasted(p)
			}
		}
	}
	MainMenu(p)
}

func BossVictory(p player.Character, b monsters.Boss) {
	Clear()
	frame1 := []string{"\033[33m" +
		"┓┏       ┏┓       ┓      ┓",
		"┣┫┓┏┏┓╋  ┃ ┏┓┏┳┓┏┓┃┏┓╋┏┓┏┫",
		"┛┗┗┛┛┗┗  ┗┛┗┛┛┗┗┣┛┗┗ ┗┗ ┗┻",
		"\033[0m",
	}
	DisplayMenu(frame1)
	Write("\n\n\nYou earned a " + b.Drop.Name + ", " + strconv.Itoa(b.Gold) + " znnys and " + strconv.Itoa(b.Exp) + " experience!\n\n")
	p.Money += b.Gold
	p.Exp += b.Exp
	player.LevelUpdate(&p)
	time.Sleep(1 * time.Second)
	Write("\n\n\nClick on anything to go back to main menu.")
	action := ""
	fmt.Scanln(&action)
	switch action {
	default:
		MainMenu(&p)
	}
}
