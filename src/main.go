package main

import (
	"fmt"
	"menu"
	"player"
	"time"
)

var Player player.Character

func main() {
	menu.Write("Hey you, yeah you, you're new in the guild right ? What's your name ? \n\n")
	name := ""
	fmt.Scanln(&name)
	Player.Name = name
	menu.Write("Ewh what a disgusting name, " + name + ", hum, do you want a beer to welcome you into the guild ?\n\n")
	menu.Write("1 - yes \n2 - no \n")
	answer1 := ""
	fmt.Scanln(&answer1)
	if answer1 == "1" {
		menu.Write("You start drinking but soon enough, you begin to stagger... the beer might have been poisonous... lol \n")
		Player.Health_point = 50
	} else if answer1 == "2" {
		menu.Write("Wait You're supposed to take it... guess we don't have the choice... \n")
		menu.Write("You got bullied and lost consciousness")
		Player.Health_point = 25
	} else if answer1 == "69" || answer1 == "420" {
		menu.Write("That's a really cool number, you still \"die\" but you're cool :D \n")
		Player.Health_point = 100
	} else {
		menu.Write("I like how you're trying to avoid dying but that's part of the game so please... die \n")
		Player.Health_point = 1
	}
	menu.Write("You wake up, somehow not dead, in a strange forest and find a \n\n")
	menu.Write("1 - GreatSword (slow but very powerful) \n2 - LongSword (fast and strong) \n3 - DualBlade (very quick but weak) \n")
	answer2 := ""
	fmt.Scanln(&answer2)
	if answer2 == "1" {
		Player.Weapon.Name = "GreatSword"
		Player.Weapon.Speed = 5
		Player.Weapon.Damage = 25
	} else if answer2 == "2" {
		Player.Weapon.Name = "LongSword"
		Player.Weapon.Speed = 15
		Player.Weapon.Damage = 15
	} else if answer2 == "3" {
		Player.Weapon.Name = "DualBlade"
		Player.Weapon.Speed = 30
		Player.Weapon.Damage = 5
	} else {
		Player.Weapon.Name = "Stick"
		Player.Weapon.Speed = 5
		Player.Weapon.Damage = 5
	}
	if Player.Name == "Kheir" || Player.Name == "Alan" || Player.Name == "Cyril" || Player.Name == "Ethan" {
		menu.Write("You just equipped the most powerful weapon in the universe, you'll get through this adventure with ease")
		Player.Weapon.Name = "Sexcalibur"
		Player.Weapon.Speed = 100
		Player.Weapon.Damage = 100
	} else {
		menu.Write("Now that you have a " + Player.Weapon.Name + " equipped, you decide to go on an adventure throughout the forest!")
	}
	Player.Level = 1
	Player.Max_health_point = 100
	Player.Money = 100
	Player.Inventory = map[string]int{"First Aid": 5}

	time.Sleep(500 * time.Millisecond)
	menu.MainMenu(Player)
}
