package main

import (
	"fmt"
	"os"
	"time"
)

type character struct {
	name             string
	weapon           string
	speed            int
	damage           int
	level            int
	max_health_point int
	health_point     int
	money            int
	inventory        map[string]int
}

func Write(sentence string) {
	for _, letter := range sentence {
		os.Stdout.WriteString((string(letter)))
		time.Sleep(100)
	}
}

var player character

func main() {
	Write("Hey you, yeah you, you're new in the guild right ? What's your name ? \n")
	name := ""
	fmt.Scanln(&name)

	player.name = name

	Write("Ewh what a disgusting name, " + name + ", hum, do you want a beer to welcome you into the guild ?\n")
	Write("1 - yes \n2 - no \n")
	answer1 := ""
	fmt.Scanln(&answer1)

	if answer1 == "1" {
		Write("You start drinking but soon enough, you begin to stagger... the beer might have been poisonous... lol")
		player.health_point = 50
	} else if answer1 == "2" {
		Write("Wait You're supposed to take it... guess we don't have the choice... \n")
		Write("You got bullied and lost consciousness")
		player.health_point = 25
	} else if answer1 == "69" || answer1 == "420" {
		Write("That's a really cool number, you still \"die\" but you're cool :D")
		player.health_point = 100
	} else {
		Write("I like how you're trying to avoid dying but that's part of the game so please... die")
		player.health_point = 1
	}

	Write("You wake up, someone not dead, in a strange forest and find a \n\n")
	Write("1 - GreatSword (slow but very powerful) \n2 - LongSword (fast and strong) \n3 - DualBlade (very quick but weak) \n")
	answer2 := ""
	fmt.Scanln(&answer2)
	if answer2 == "1" {
		player.weapon = "GreatSword"
		player.speed = 5
		player.damage = 25
	} else if answer2 == "2" {
		player.weapon = "LongSword"
		player.speed = 15
		player.damage = 15
	} else if answer2 == "3" {
		player.weapon = "DualBlade"
		player.speed = 30
		player.damage = 5
	} else {
		player.weapon = "Stick"
		player.speed = 5
		player.damage = 5
	}
	Write("Now that you have a " + player.weapon + " equipped, you decide to go on an adventure throughout the forest!")

	player.level = 1
	player.max_health_point = 100
	player.money = 100
	player.inventory = map[string]int{"First Aid": 3}
}
