package main

import (
	"fmt"
	"menu"
	"player"
	"time"
)

var p1 player.Character

func main() {
	p1.Stack = 5
	p1.Affliction = "None"
	menu.Clear()
	menu.Write("Hey you, yeah you, you're new in the guild right ? What's your name ? \n\n")
	name := ""
	fmt.Scanln(&name)
	p1.Name = Capitalize(name)
	menu.Clear()
	menu.Write("Ewh what a disgusting name, " + p1.Name + ", hum, do you want a beer to welcome you into the guild ?\n\n")
	menu.Write("1 - yes \n2 - no \n")
	answer1 := ""
	fmt.Scanln(&answer1)
	menu.Clear()
	if answer1 == "1" {
		menu.Write("You start drinking but soon enough, you begin to stagger... the beer might have been poisonous... lol \n")
		p1.Health_point = 50
	} else if answer1 == "2" {
		menu.Write("Wait You're supposed to take it... guess we don't have the choice... \n")
		menu.Write("You got bullied and lost consciousness")
		p1.Health_point = 25
	} else if answer1 == "69" || answer1 == "420" {
		menu.Write("That's a really cool number, you still \"die\" but you're cool :D \n")
		p1.Health_point = 100
	} else {
		menu.Write("I like how you're trying to avoid dying but that's part of the game so please... die \n")
		p1.Health_point = 1
	}
	time.Sleep(3 * time.Second)
	menu.Clear()
	menu.Write("You wake up, somehow not dead, in a strange forest and find a \n\n")
	menu.Write("1 - GreatSword (slow but very powerful) \n2 - LongSword (fast and strong) \n3 - DualBlade (very quick but weak) \n")
	answer2 := ""
	fmt.Scanln(&answer2)
	if answer2 == "1" {
		p1.Weapon.Name = "Iron"
		p1.Weapon.Type = "GreatSword"
		p1.Weapon.Speed = 5
		p1.Weapon.Damage = 25
	} else if answer2 == "2" {
		p1.Weapon.Name = "Iron"
		p1.Weapon.Type = "LongSword"
		p1.Weapon.Speed = 15
		p1.Weapon.Damage = 15
	} else if answer2 == "3" {
		p1.Weapon.Name = "Iron"
		p1.Weapon.Type = "DualBlades"
		p1.Weapon.Speed = 30
		p1.Weapon.Damage = 5
	} else {
		p1.Weapon.Name = "Wooden"
		p1.Weapon.Type = "Stick"
		p1.Weapon.Speed = 5
		p1.Weapon.Damage = 5
	}
	menu.Clear()
	if p1.Name == "Kheir" || p1.Name == "Alan" || p1.Name == "Cyril" || p1.Name == "Ethan" || p1.Name == "Fhaaab" || p1.Name == "Plcuf" {
		menu.Write("You just equipped the most powerful weapon in the universe, you'll get through this adventure with ease")
		p1.Weapon.Name = "Ultimate Sword"
		p1.Weapon.Type = "Sexcalibur"
		p1.Weapon.Speed = 100
		p1.Weapon.Damage = 100
	} else {
		menu.Write("Now that you have a " + p1.Weapon.Name + " " + p1.Weapon.Type + " equipped, you decide to go on an adventure throughout the forest!")
	}
	p1.Level = 1
	p1.Max_health_point = 100
	p1.Money = 100
	p1.Inventory = map[player.Item]int{player.First_Aid: 5}
	p1.Armor.Helm = "Leather"
	p1.Armor.Chest = "Leather"
	p1.Armor.Arms = "Leather"
	p1.Armor.Waist = "Leather"
	p1.Armor.Legs = "Leather"
	time.Sleep(3 * time.Second)
	menu.MainMenu(p1)
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
