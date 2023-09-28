package menu

import (
	"fmt"
	"math/rand"
	"red/global/monsters"
	"red/global/player"
	"strconv"
	"time"
)

func BossBattle(p *player.Character, b *monsters.Boss) {
	playerHP := &p.Health_point
	monsterHP := b.PV
	stamina := p.Stamina_max
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
		if stamina < 100 {
			stamina += 10
			if stamina > p.Stamina_max {
				stamina = p.Stamina_max
			}
		}
		if speedp >= 1000 {
			speedp -= 1000
			battleMenu := []string{
				"Battle : \n\n\n",
				"   " + b.Name + "   " + "\033[31m" + strconv.Itoa(monsterHP) + "/" + strconv.Itoa(b.PV) + "\033[0m" + "\n\n",
				" ↳ " + p.Name + "   " + "\033[31m" + strconv.Itoa(*playerHP) + "/" + strconv.Itoa(p.Max_health_point) + "\033[0m" + "   " + "\033[36m" + strconv.Itoa(stamina) + "/" + strconv.Itoa(p.Stamina_max) + "\033[0m"}
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
					monsterHP -= p.Weapon.Skills[0].Use(p)
					stamina -= p.Weapon.Skills[0].Cost
					Write("Skill used !")
					time.Sleep(time.Second)
					if monsterHP <= 0 {
						VictoryBoss(p, *b)
						hunting = false
					}
				case "2":
					monsterHP -= p.Weapon.Skills[1].Use(p)
					stamina -= p.Weapon.Skills[1].Cost
					Write("Skill used !")
					time.Sleep(time.Second)
					if monsterHP <= 0 {
						VictoryBoss(p, *b)
						hunting = false
					}
				case "3":
					monsterHP -= p.Weapon.Skills[2].Use(p)
					stamina -= p.Weapon.Skills[2].Cost
					Write("Skill used !")
					time.Sleep(time.Second)
					if monsterHP <= 0 {
						VictoryBoss(p, *b)
						hunting = false
					}
				}
			case "2":
				Clear()
				inventory := []player.Item{}
				for object := range p.Inventory {
					if object.Usable {
						inventory = append(inventory, object)
					}
				}
				navigateMenu := []string{"Select your object\n\n"}
				actionCount := 0
				for i := 0; i < len(inventory); i++ {
					if inventory[i].Usable {
						actionCount++
						navigateMenu = append(navigateMenu, strconv.Itoa(actionCount)+" - "+inventory[i].Name)
					}
				}
				Clear()
				DisplayMenu(navigateMenu)
				Write("\n\n0 - Back")
				navigate := ""
				fmt.Scanln(&navigate)
				response, _ := strconv.Atoi(navigate)
				if response > 0 && response <= len(inventory) {
					inventory[response-1].Use(p)
				}
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
				" ↱ " + b.Name + "   " + strconv.Itoa(monsterHP) + "/" + strconv.Itoa(b.PV) + "\n\n",
				"   " + p.Name + "   " + "\033[31m" + strconv.Itoa(*playerHP) + "/" + strconv.Itoa(p.Max_health_point) + "\033[0m" + "   " + "\033[36m" + strconv.Itoa(stamina) + "/" + strconv.Itoa(p.Stamina_max) + "\033[0m"}
			Clear()
			DisplayMenu(battleMenu)
			fmt.Println("Monster's turn !")
			*playerHP -= b.Damage
			Write(b.Name + " attacked !")
			time.Sleep(time.Second)
			if *playerHP <= 0 {
				hunting = false
				Wasted(p)
			}
		}
	}
	MainMenu(p)
}

func VictoryBoss(p *player.Character, b monsters.Boss) {
	Clear()
	frame1 := []string{"\033[33m" +
		"┓┏       ┏┓       ┓      ┓",
		"┣┫┓┏┏┓╋  ┃ ┏┓┏┳┓┏┓┃┏┓╋┏┓┏┫",
		"┛┗┗┛┛┗┗  ┗┛┗┛┛┗┗┣┛┗┗ ┗┗ ┗┻",
		"\033[0m",
	}
	DisplayMenu(frame1)
	Write("\n\n\nYou earned a " + b.Drop.Name + ", " + strconv.Itoa(b.Gold) + " zennys and " + strconv.Itoa(b.Exp) + " experience!\n\n")
	chance := rand.Int31n(100)
	if chance <= int32(b.Drop_chances*100.0) {
		p.Inventory[b.Drop] += 1
	}
	p.Money += b.Gold
	p.Exp += b.Exp
	player.LevelUpdate(p)
	time.Sleep(1 * time.Second)
	Write("\n\n\nClick on anything to go back to main menu.")
	action := ""
	fmt.Scanln(&action)
	switch action {
	default:
		if b.Name == "Great Jagras" && p.Quest == 0 {
			p.Quest += 1
		} else if b.Name == "Barroth" && p.Quest == 1 {
			p.Quest += 1
		} else if b.Name == "Rathalos" && p.Quest == 2 {
			p.Quest += 1
		} else if b.Name == "Nergigante" && p.Quest == 3 {
			p.Quest += 1
		} else if b.Name == "Xeno'Jiiva" && p.Quest == 4 {
			p.Quest += 1
		}
		MainMenu(p)
	}
}
