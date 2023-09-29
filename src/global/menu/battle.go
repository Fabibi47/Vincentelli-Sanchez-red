package menu

import (
	"fmt"
	"math/rand"
	"red/global/monsters"
	"red/global/player"

	"strconv"
	"time"
)

func ChoseZone(player *player.Character) {
	Clear()
	ChoiceMenu := []string{
		"\033[91m" + "   __    ____  ____    __    ___     ",
		"  /__\\  (  _ \\( ___)  /__\\  / __)  ()",
		" /(__)\\  )   / )__)  /(__)\\ \\__ \\    ",
		"(__)(__)(_)\\_)(____)(__)(__)(___/  ()\n\n",
		"\033[90m" + "	1 - Forest\n",
		"	2 - Desert\n",
		"	3 - Swamps\n",
		"	4 - Ancient's Land\n" + "\033[0m",
		"0 - Back\n\n"}
	DisplayMenu(ChoiceMenu)
	action := ""
	fmt.Scanln(&action)
	switch action {
	case "1":
		Hunt(player, GetMonster("Forest"), "Forest")
	case "2":
		Hunt(player, GetMonster("Desert"), "Desert")
	case "3":
		Hunt(player, GetMonster("Swamp"), "Swamp")
	case "4":
		Hunt(player, GetMonster("Ancient Land"), "Ancient Land")
	default:
		MainMenu(player)
	}
}

func GetMonster(zone string) monsters.Monster {
	var possible []monsters.Monster
	if zone == "Forest" {
		possible = []monsters.Monster{monsters.Jagras, monsters.Mernos, monsters.Aptonoth, monsters.Kestodon}
	} else if zone == "Desert" {
		possible = []monsters.Monster{monsters.Apceros, monsters.Noios, monsters.Dellex}
	} else if zone == "Swamp" {
		possible = []monsters.Monster{monsters.Shamos, monsters.Raphinos, monsters.Girros}
	} else if zone == "Ancient Land" {
		possible = []monsters.Monster{monsters.Barnos, monsters.Dodogama, monsters.Uragaan}
	}
	return possible[rand.Intn(len(possible))]
}

func Hunt(p *player.Character, monster monsters.Monster, zone string) {
	playerHP := &p.Health_point
	monsterHP := monster.PV
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
		speedm += monster.Speed
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
				"   " + monster.Name + "   " + "\033[31m" + strconv.Itoa(monsterHP) + "/" + strconv.Itoa(monster.PV) + "\033[0m" + "\n\n",
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
						Victory(p, monster, zone)
						hunting = false
					}
					if p.Weapon.Effect == "Poison" {
						monster.Affliction = "Poison"
					}
				case "2":
					monsterHP -= p.Weapon.Skills[1].Use(p)
					stamina -= p.Weapon.Skills[1].Cost
					Write("Skill used !")
					time.Sleep(time.Second)
					if monsterHP <= 0 {
						Victory(p, monster, zone)
						hunting = false
					}
					if p.Weapon.Effect == "Poison" {
						monster.Affliction = "Poison"
					}
				case "3":
					monsterHP -= p.Weapon.Skills[2].Use(p)
					stamina -= p.Weapon.Skills[2].Cost
					Write("Skill used !")
					time.Sleep(time.Second)
					if monsterHP <= 0 {
						Victory(p, monster, zone)
						hunting = false
					}
					if p.Weapon.Effect == "Poison" {
						monster.Affliction = "Poison"
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
				" ↱ " + monster.Name + "   " + strconv.Itoa(monsterHP) + "/" + strconv.Itoa(monster.PV) + "\n\n",
				"   " + p.Name + "   " + "\033[31m" + strconv.Itoa(*playerHP) + "/" + strconv.Itoa(p.Max_health_point) + "\033[0m" + "   " + "\033[36m" + strconv.Itoa(stamina) + "/" + strconv.Itoa(p.Stamina_max) + "\033[0m"}
			Clear()
			DisplayMenu(battleMenu)
			fmt.Println("Monster's turn !")
			*playerHP -= monster.Damage
			Write(monster.Name + " attacked !")
			time.Sleep(time.Second)
			if *playerHP <= 0 {
				hunting = false
				Wasted(p)
			}
		}
		if monster.Affliction == "Poison" {
			monsterHP -= int(float32(monster.PV) * 0.05)
			Write("Monster poisonned !")
			time.Sleep(time.Second)
		}
	}
	MainMenu(p)
}

func GetSkills(p *player.Character) []string {
	skillList := []string{}
	for i, skill := range p.Weapon.Skills {
		skillList = append(skillList, "\n"+strconv.Itoa(i+1)+" - "+skill.Name)
	}
	skillList = append(skillList, "\n0 - Back")
	return skillList
}
