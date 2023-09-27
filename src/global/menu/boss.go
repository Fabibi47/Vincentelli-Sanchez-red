package menu

import (
	"fmt"
	"math/rand"
	"monsters"
	"player"
	"strconv"
	"time"
)

func Boss(p player.Character, b monsters.Boss) {
	p.Affliction = "None"
	playerHP := p.Health_point
	bossHP := b.PV
	savebossHP := b.PV
	hunting := true
	playerAction := []string{"1 - Attack\n", "2 - Objects\n", "3 - Run"}
	turns := GetTurnsBoss(p, b)
	for i := 0; i < len(turns) && hunting; i++ {
		battleMenu := []string{p.Name, "PV :" + strconv.Itoa(bossHP), p.Name, "PV : " + strconv.Itoa(playerHP) + "/" + strconv.Itoa(p.Max_health_point)}
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
				DisplayMenu(GetSkills(p))
				attack := ""
				fmt.Scanln(&attack)
				switch attack {
				case "1":
					bossHP -= p.Weapon.Skills[0].Use(&p)
					Write("Skill used !")
					time.Sleep(time.Second)
					b.Affliction = p.Weapon.Skills[0].Effect
					if bossHP <= 0 {
						VictoryBoss(&p, b)
						hunting = false
					}
				case "2":
					bossHP -= p.Weapon.Skills[1].Use(&p)
					Write("Skill used !")
					time.Sleep(time.Second)
					b.Affliction = p.Weapon.Skills[0].Effect
					if bossHP <= 0 {
						VictoryBoss(&p, b)
						hunting = false
					}
				case "3":
					bossHP -= p.Weapon.Skills[2].Use(&p)
					Write("Skill used !")
					time.Sleep(time.Second)
					b.Affliction = p.Weapon.Skills[0].Effect
					if bossHP <= 0 {
						VictoryBoss(&p, b)
						hunting = false
					}
				}
			case "2":
				Clear()
				DisplayMenu(battleMenu)
				Items(&p)
			case "3":
				Clear()
				Write("You ran away... Loser.")
				time.Sleep(3 * time.Second)
				hunting = false
			}
		} else {
			fmt.Println("Monster's turn !")
			chance := rand.Int()
			if chance%5 == 0 {
				playerHP -= b.Attack.Use(&b)
				Write("The Boss used his Skill !")
				if b.Attack.Effect == "Burn" {
					p.Affliction = "Burn"
					Write("You're burning !")
				}
			}
			Write(b.Name + " attacked !")
			time.Sleep(time.Second)
			if playerHP <= 0 {
				hunting = false
				Wasted(&p)
			}
		}
		if i == len(turns)-1 {
			if b.Affliction == "Poison" {
				bossHP -= int(float64(savebossHP) * 0.05)
				Write("Ennemy poisonned.")
				time.Sleep(time.Second)
			}
			if p.Affliction == "Burn" {
				playerHP -= 5
			}
			i = 0
		}
	}
	MainMenu(p)
}

func GetTurnsBoss(player player.Character, boss monsters.Boss) []string {
	turns := []string{}
	if player.Weapon.Speed > boss.Speed {
		nb := player.Weapon.Speed / boss.Speed
		for ; nb > 0; nb-- {
			turns = append(turns, "player")
		}
		turns = append(turns, "Boss")
	} else {
		nb := boss.Speed / player.Weapon.Speed
		for ; nb > 0; nb-- {
			turns = append(turns, "Boss")
		}
		turns = append(turns, "player")
	}
	return turns
}

func VictoryBoss(p *player.Character, b monsters.Boss) {
	Clear()
	Write("Hunt Completed !\n\n\n\n")
	p.Inventory[b.Drop] += 1
	Write("You earned a" + b.Drop.Name + "!")
	time.Sleep(3 * time.Second)
	Write("0 - Back\n")
	action := ""
	fmt.Scanln(&action)
	switch action {
	default:
		MainMenu(*p)
	}
}
