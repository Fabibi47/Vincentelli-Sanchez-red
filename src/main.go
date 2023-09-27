package main

import (
	"colors"
	"fmt"
	"menu"
	"player"
	"time"
)

var p1 player.Character

func main() {
<<<<<<< HEAD
	p1.Stack = 5
	p1.Affliction = "None"
=======
	colors.Init()
>>>>>>> f1860076939a29938ee7eebbced9925ae3096641
	menu.Clear()
	menu.Write1("Hey you, yeah you, you're new in the guild right ? What's your name ? \n\n")
	name := ""
	fmt.Scanln(&name)
	p1.Name = Capitalize(name)
	menu.Clear()
	menu.Write1("Ewh what a disgusting name, " + p1.Name + ", hum, do you want a beer to welcome you into the guild ?\n\n")
	menu.Write1("1 - yes \n2 - no \n")
	answer1 := ""
	fmt.Scanln(&answer1)
	menu.Clear()
	if answer1 == "1" {
		menu.Write1("You start drinking but soon enough, you begin to stagger... the beer might have been poisonous... lol \n")
		p1.Health_point = 50
	} else if answer1 == "2" {
		menu.Write1("Wait You're supposed to take it... guess we don't have the choice... \n")
		menu.Write1("You got bullied and lost consciousness")
		p1.Health_point = 25
	} else if answer1 == "69" || answer1 == "420" {
		menu.Write1("That's a really cool number, you still \"die\" but you're cool :D \n")
		p1.Health_point = 100
	} else {
		menu.Write1("I like how you're trying to avoid dying but that's part of the game so please... die \n")
		p1.Health_point = 1
	}
	time.Sleep(3 * time.Second)
	menu.Clear()
	menu.Write1("You wake up, somehow not dead, in a strange forest and find a \n\n")
	menu.Write1("1 - GreatSword (slow but very powerful) \n2 - LongSword (fast and strong) \n3 - DualBlade (very quick but weak) \n")
	answer2 := ""
	fmt.Scanln(&answer2)
	if answer2 == "1" {
		p1.Weapon = player.Weapon{Name: "Iron", Damage: 25, Speed: 5, Level: 1, Type: "GreatSword"}
	} else if answer2 == "2" {
		p1.Weapon = player.Weapon{Name: "Iron", Damage: 15, Speed: 15, Level: 1, Type: "LongSword"}
	} else if answer2 == "3" {
		p1.Weapon = player.Weapon{Name: "Iron", Damage: 10, Speed: 25, Level: 1, Type: "DualBlades"}
	} else {
		p1.Weapon = player.Weapon{Name: "Wooden", Damage: 5, Speed: 5, Level: 1, Type: "Stick"}
	}
	menu.Clear()
	if p1.Name == "Kheir" || p1.Name == "Alan" || p1.Name == "Cyril" || p1.Name == "Ethan" || p1.Name == "Fhaaab" || p1.Name == "Plcuf" {
		menu.Write1("You just equipped the most powerful weapon in the universe, you'll get through this adventure with ease")
		p1.Weapon.Name = "Ultimate Sword"
		p1.Weapon.Type = "Sexcalibur"
		p1.Weapon.Speed = 100
		p1.Weapon.Damage = 100
	} else {
		menu.Write1("Now that you have a " + p1.Weapon.Name + " " + p1.Weapon.Type + " equipped, you decide to go on an adventure throughout the forest!")
	}
	p1.Level = 1
	p1.Base_hp = 100
	p1.Money = 10000
	p1.Inventory = map[player.Item]int{player.First_Aid: 5, player.Iron_ore: 5, player.Machalite_ore: 5, player.Dragonite_ore: 5, player.Carbalite_ore: 5, player.Fucium_ore: 5, player.Eltalite_ore: 5}
	p1.Armor.Helm = "Leather"
	p1.Armor.Chest = "Leather"
	p1.Armor.Arms = "Leather"
	p1.Armor.Waist = "Leather"
	p1.Armor.Legs = "Leather"
	p1.Max_health_point = player.HpUpdate(&p1.Armor, &p1)
	player.SkillDetection(&p1.Weapon)
	time.Sleep(3 * time.Second)
	menu.MainMenu(&p1)
}

func Capitalize(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' && i == 0 {
			result += string(rune(s[i]) - 32)
		} else if s[i] >= 'a' && s[i] <= 'z' && (s[i-1] < 'A' || s[i-1] > 'Z') && (s[i-1] < 'a' || s[i-1] > 'z') && (s[i-1] < '0' || s[i-1] > '9') {
			result += string(rune(s[i]) - 32)
		} else if s[i] >= 'A' && s[i] <= 'Z' && i == 0 {
			result += string(rune(s[i]))
		} else if s[i] >= 'A' && s[i] <= 'Z' && (s[i-1] >= 'A' && s[i-1] <= 'Z' || s[i-1] >= 'a' && s[i-1] <= 'z' || s[i-1] >= '0' && s[i-1] <= '9') {
			result += string(rune(s[i]) + 32)
		} else {
			result += string(rune(s[i]))
		}
	}
	return result
}
