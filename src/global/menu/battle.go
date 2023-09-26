package menu

import (
	"fmt"
	"math/rand"
	"monsters"
	"player"
	"strconv"
	"time"
)

func ChoseZone(player player.Character) {
	Clear()
	Write("Which region do you want to go hunting?")
	ChoiceMenu := []string{"1 - Forest\n", "2 - Desert\n", "3 - Swamps\n\n", "0 - Back"}
	DisplayMenu(ChoiceMenu)
	action := ""
	fmt.Scanln(&action)
	switch action {
	case "1":
		Hunt(player, GetMonster("Forest"), "Forest")
	case "2":
		Hunt(player, GetMonster("Desert"), "Desert")
	case "3":
		Hunt(player, GetMonster("Swamp"), "Desert")
	default:
		MainMenu(&player)
	}
}

func GetMonster(zone string) monsters.Monster {
	var possible []monsters.Monster
	if zone == "Forest" {
		possible = []monsters.Monster{monsters.Jagras, monsters.Mernos}
	} else if zone == "Desert" {
		possible = []monsters.Monster{monsters.Apceros, monsters.Noios}
	} else if zone == "Swamp" {
		possible = []monsters.Monster{monsters.Shamos, monsters.Raphinos}
	}
	return possible[rand.Intn(len(possible))]
}

func Hunt(player player.Character, monster monsters.Monster, zone string) {
	playerHP := &player.Health_point
	monsterHP := &monster.PV
	saveMonsterHP := monster.PV
	hunting := true
	playerAction := []string{"1 - Attack\n", "2 - Objects\n", "3 - Run"}
	turns := GetTurns(player, monster)
	for i := 0; i < len(turns) && hunting; i++ {
		battleMenu := []string{monster.Name, "PV :" + strconv.Itoa(*monsterHP), player.Name, "PV : " + strconv.Itoa(*playerHP) + "/" + strconv.Itoa(player.Max_health_point)}
		Clear()
		DisplayMenu(battleMenu)
		if turns[i] == "player" {
			fmt.Println("Your turn.")
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
					*monsterHP -= player.Weapon.Skills[0].Use(&player)
					Write("Skill used !")
					time.Sleep(time.Second)
					monster.Affliction = player.Weapon.Skills[0].Effect
					if *monsterHP <= 0 {
						Victory(&player, monster, zone)
						hunting = false
					}
				case "2":
					*monsterHP -= player.Weapon.Skills[1].Use(&player)
					Write("Skill used !")
					time.Sleep(time.Second)
					monster.Affliction = player.Weapon.Skills[0].Effect
					if *monsterHP <= 0 {
						Victory(&player, monster, zone)
						hunting = false
					}
				case "3":
					*monsterHP -= player.Weapon.Skills[2].Use(&player)
					Write("Skill used !")
					time.Sleep(time.Second)
					monster.Affliction = player.Weapon.Skills[0].Effect
					if *monsterHP <= 0 {
						Victory(&player, monster, zone)
						hunting = false
					}
				}
			case "2":
				Clear()
				DisplayMenu(battleMenu)
				Items(&player)
			case "3":
				Clear()
				Write("You ran away... Loser.")
				time.Sleep(3 * time.Second)
				hunting = false
			}
		} else {
			fmt.Println("Monster's turn !")
			*playerHP -= monster.Damage
			Write(monster.Name + " attacked !")
			time.Sleep(time.Second)
			if *playerHP <= 0 {
				Write("You lost...")
				time.Sleep(time.Second * 3)
				hunting = false
				MainMenu(&player)
			}
		}
		if i == len(turns)-1 {
			if monster.Affliction == "Poison" {
				*monsterHP -= int(float64(saveMonsterHP) * 0.05)
				Write("Ennemy poisonned.")
				time.Sleep(time.Second)
			}
			i = 0
		}
	}
	MainMenu(&player)
}

func GetTurns(player player.Character, monster monsters.Monster) []string {
	turns := []string{}
	if player.Weapon.Speed > monster.Speed {
		nb := player.Weapon.Speed / monster.Speed
		for ; nb > 0; nb-- {
			turns = append(turns, "player")
		}
		turns = append(turns, "monster")
	} else {
		nb := monster.Speed / player.Weapon.Speed
		for ; nb > 0; nb-- {
			turns = append(turns, "monster")
		}
		turns = append(turns, "player")
	}
	return turns
}

func GetSkills(p player.Character) []string {
	skillList := []string{}
	for i, skill := range p.Weapon.Skills {
		skillList = append(skillList, strconv.Itoa(i+1)+" - "+skill.Name+"\n")
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

func Victory(p *player.Character, m monsters.Monster, zone string) {
	Clear()
	Write("Hunt Completed !\n\n\n\n")
	p.Money += m.Drop
	Write("You earned 10 Zennys !\n\n")
	time.Sleep(3 * time.Second)
	Write("1 - Next monster\n")
	Write("0 - Back\n")
	action := ""
	fmt.Scanln(&action)
	switch action {
	case "1":
		Hunt(*p, GetMonster(zone), zone)
	default:
		MainMenu(p)
	}
}
