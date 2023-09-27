package menu

import (
	"fmt"
	"math/rand"
	"monsters"
	"player"
	"strconv"
	"time"
)

func ChoseZone(player *player.Character) {
	Clear()
	ChoiceMenu := []string{
		"In which area do you want to hunt ?\n\n",
		"	1 - Forest\n",
		"	2 - Desert\n",
		"	3 - Swamps\n",
		"	4 - Ancient's Land\n",
		"	5 - UnderWorld\n",
		"	6 - Sacred Lands\n\n",
		"0 - Back"}
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

func Hunt(player *player.Character, monster monsters.Monster, zone string) {
	playerHP := &player.Health_point
	monsterHP := monster.PV
	hunting := true
	playerAction := []string{
		"\n\n	1 - Attack",
		"	2 - Objects",
		"	3 - Run\n\n"}
	speedp := 0
	speedm := 0
	for hunting {
		speedp += player.Weapon.Speed
		speedm += monster.Speed
		if speedp >= 1000 {
			speedp -= 1000
			battleMenu := []string{
				"Battle : \n\n\n",
				"   " + monster.Name + "   " + strconv.Itoa(monsterHP) + "/" + strconv.Itoa(monster.PV) + "\n\n",
				" ↳ " + player.Name + "   " + strconv.Itoa(*playerHP) + "/" + strconv.Itoa(player.Max_health_point)}
			Clear()
			DisplayMenu(battleMenu)
			DisplayMenu(playerAction)
			action := ""
			fmt.Scanln(&action)
			switch action {
			case "1":
				Clear()
				DisplayMenu(battleMenu)
				DisplayMenu(GetSkills(player))
				attack := ""
				fmt.Scanln(&attack)
				switch attack {
				case "1":
					monsterHP -= player.Weapon.Skills[0].Use(player)
					Write("Skill used !")
					time.Sleep(time.Second)
					if monsterHP <= 0 {
						Victory(player, monster, zone)
						hunting = false
					}
				case "2":
					monsterHP -= player.Weapon.Skills[1].Use(player)
					Write("Skill used !")
					time.Sleep(time.Second)
					if monsterHP <= 0 {
						Victory(player, monster, zone)
						hunting = false
					}
				case "3":
					monsterHP -= player.Weapon.Skills[2].Use(player)
					Write("Skill used !")
					time.Sleep(time.Second)
					if monsterHP <= 0 {
						Victory(player, monster, zone)
						hunting = false
					}
				}
			case "2":
				Clear()
				DisplayMenu(battleMenu)
				Items(player)
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
				"   " + player.Name + "   " + strconv.Itoa(*playerHP) + "/" + strconv.Itoa(player.Max_health_point)}
			Clear()
			DisplayMenu(battleMenu)
			fmt.Println("Monster's turn !")
			*playerHP -= monster.Damage
			Write(monster.Name + " attacked !")
			time.Sleep(time.Second)
			if *playerHP <= 0 {
				hunting = false
				Wasted(player)
			}
		}
	}
	MainMenu(player)
}

func GetSkills(p *player.Character) []string {
	skillList := []string{}
	for i, skill := range p.Weapon.Skills {
		skillList = append(skillList, "\n\n	"+strconv.Itoa(i+1)+" - "+skill.Name+"\n")
	}
	skillList = append(skillList, "0 - Back")
	return skillList
}

func Items(p *player.Character) {
	items := []player.Item{}
	for item := range p.Inventory {
		if item.Usable {
			items = append(items, item)
		}
	}
	selectionSlice := []string{}
	for j, i := range items {
		selectionSlice = append(selectionSlice, strconv.Itoa(j+1)+" - "+i.Name)
	}
	DisplayMenu(selectionSlice)
	Write("0 - Back")
	action := ""
	fmt.Scanln(&action)
	if action > strconv.Itoa(0) && action <= strconv.Itoa(len(selectionSlice)) {
		item, _ := strconv.Atoi(action)
		item--
		items[item].Use(p)
	}
}
