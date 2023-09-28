package menu

import (
	"bufio"
	"os"
	"os/exec"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)

func DisplayMenu(menu []string) {
	for _, s := range menu {
		Write(s)
	}
}

func Write1(s string) {
	for _, c := range s {
		os.Stdout.WriteString(string(c))
		time.Sleep(30 * time.Millisecond)
	}
	os.Stdout.WriteString("\n")
}

func Write(s string) {
	for _, c := range s {
		os.Stdout.WriteString(string(c))
	}
	os.Stdout.WriteString("\n")
}

func Clear() { //ne marche que sous Windows
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}
