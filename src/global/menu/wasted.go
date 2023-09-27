package menu

import (
	"fmt"
	"player"
	"time"
)

func freeze() {
	time.Sleep(50 * time.Millisecond)
	Clear()
}

func Wasted(p *player.Character) {
	Clear()
	frame1 := []string{
		"█     █                                             ",
		" █  █  █                                            ",
		" █  █  █                                            ",
		" █  █  █                                            ",
		"  ██ ██                                             ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
	}
	DisplayMenu(frame1)
	freeze()
	frame2 := []string{
		"█     █  ▄▄▄                                        ",
		" █  █  █  ████▄                                     ",
		" █  █  █  ██  ▀█▄                                   ",
		" █  █  █  ██▄▄▄▄██                                  ",
		"  ██ ██    █    ██                                  ",
		"                 █                                  ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
	}
	DisplayMenu(frame2)
	freeze()
	frame3 := []string{
		"█     █  ▄▄▄        ██████                          ",
		" █  █  █  ████▄     ██                              ",
		" █  █  █  ██  ▀█▄     ██▄                           ",
		" █  █  █  ██▄▄▄▄██       ██                         ",
		"  ██ ██    █    ██  ██████                          ",
		"                 █                                  ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
	}
	DisplayMenu(frame3)
	freeze()
	frame4 := []string{
		"█     █  ▄▄▄        ██████ ▄▄▄█████                 ",
		" █  █  █  ████▄     ██         ██                   ",
		" █  █  █  ██  ▀█▄     ██▄      ██                   ",
		" █  █  █  ██▄▄▄▄██       ██    ██                   ",
		"  ██ ██    █    ██  ██████     ██                   ",
		"                 █                                  ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
	}
	DisplayMenu(frame4)
	freeze()
	frame5 := []string{
		"█     █  ▄▄▄        ██████ ▄▄▄█████  █████          ",
		" █  █  █  ████▄     ██         ██     █             ",
		" █  █  █  ██  ▀█▄     ██▄      ██     ███           ",
		" █  █  █  ██▄▄▄▄██       ██    ██      █            ",
		"  ██ ██    █    ██  ██████     ██      ████         ",
		"                 █                                  ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
	}
	DisplayMenu(frame5)
	freeze()
	frame6 := []string{
		"█     █  ▄▄▄        ██████ ▄▄▄█████  █████  █████▄  ",
		" █  █  █  ████▄     ██         ██     █      ██  ██▌",
		" █  █  █  ██  ▀█▄     ██▄      ██     ███    ██   █▌",
		" █  █  █  ██▄▄▄▄██       ██    ██      █      █    ▌",
		"  ██ ██    █    ██  ██████     ██      ████   █████▌",
		"                 █                                  ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
	}
	DisplayMenu(frame6)
	time.Sleep(250 * time.Millisecond)
	Clear()
	frame7 := []string{
		"█     █  ▄▄▄        ██████ ▄▄▄█████" + "\033[31m" + "▓▓" + "\033[0m" + "█████ " + "\033[31m" + "▓" + "\033[0m" + "█████▄  ",
		"\033[31m" + "▓" + "\033[0m" + "█  █  █ " + "\033[31m" + "▒" + "\033[0m" + "████▄    " + "\033[31m" + "▒" + "\033[0m" + "██    " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "  ██  " + "\033[31m" + "▓" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "█     " + "\033[31m" + "▒" + "\033[0m" + "██  ██▌",
		"\033[31m" + "▒" + "\033[0m" + "█  █  █  ██  ▀█▄    " + "\033[31m" + "▓" + "\033[0m" + "██▄     " + "\033[31m" + "▓" + "\033[0m" + "██    " + "\033[31m" + "▒" + "\033[0m" + "███    ██   █▌",
		" █  █  █  ██▄▄▄▄██   " + "\033[31m" + "▒" + "\033[0m" + "   ██" + "\033[31m" + "▒" + "\033[0m" + "  " + "\033[31m" + "▓" + "\033[0m" + "██" + "\033[31m" + "▓" + "\033[0m" + "   " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "█     " + "\033[31m" + "▓" + "\033[0m" + "█    ▌",
		"  ██ ██" + "\033[31m" + "▓" + "\033[0m" + "  " + "\033[31m" + "▓" + "\033[0m" + "█   " + "\033[31m" + "▓" + "\033[0m" + "██  ██████" + "\033[31m" + "▒" + "\033[0m" + "    ██     " + "\033[31m" + "▒" + "\033[0m" + "████  " + "\033[31m" + "▒" + "\033[0m" + "█████▌ ",
		"  " + "\033[31m" + "▓" + "\033[0m" + "            " + "\033[31m" + "▓" + "\033[0m" + " █    " + "\033[31m" + "▓" + "\033[0m" + "  " + "\033[31m" + "▒" + "\033[0m" + "              " + "\033[31m" + "▒" + "\033[0m" + "      " + "\033[31m" + "▓" + "\033[0m" + "    ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
		"                                                    ",
	}
	DisplayMenu(frame7)
	time.Sleep(250 * time.Millisecond)
	Clear()
	frame8 := []string{
		"█     █  ▄▄▄        ██████ ▄▄▄█████" + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "█████ " + "\033[31m" + "▓" + "\033[0m" + "█████▄  ",
		"" + "\033[31m" + "▓" + "\033[0m" + "█  █  █ " + "\033[31m" + "▒" + "\033[0m" + "████▄    " + "\033[31m" + "▒" + "\033[0m" + "██    " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "  ██" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "█     " + "\033[31m" + "▒" + "\033[0m" + "██  ██▌",
		"" + "\033[31m" + "▓" + "\033[0m" + "█  █  █ " + "\033[31m" + "▒" + "\033[0m" + "██  ▀█▄    " + "\033[31m" + "▓" + "\033[0m" + "██▄   " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "██  " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "███    ██   █▌",
		"" + "\033[31m" + "░" + "\033[0m" + "█  █  █  ██▄▄▄▄██   " + "\033[31m" + "▒" + "\033[0m" + "   ██" + "\033[31m" + "▒" + "\033[0m" + "  " + "\033[31m" + "▓" + "\033[0m" + "██" + "\033[31m" + "▓" + "\033[0m" + "   " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "█     " + "\033[31m" + "▓" + "\033[0m" + "█    ▌",
		"  ██" + "\033[31m" + "▒" + "\033[0m" + "██" + "\033[31m" + "▓" + "\033[0m" + "  " + "\033[31m" + "▓" + "\033[0m" + "█   " + "\033[31m" + "▓" + "\033[0m" + "██" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "██████" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "  " + "\033[31m" + "▒" + "\033[0m" + "██" + "\033[31m" + "▒" + "\033[0m" + "    " + "\033[31m" + "▒" + "\033[0m" + "████" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "█████▌ ",
		"  " + "\033[31m" + "▓" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "   " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "   " + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "█ " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "    " + "\033[31m" + "▒" + "\033[0m" + "         " + "\033[31m" + "▒" + "\033[0m" + "    " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "  " + "\033[31m" + "▒" + "\033[0m" + " ",
		"         " + "\033[31m" + "▒" + "\033[0m" + "   " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "     " + "\033[31m" + "▒" + "\033[0m" + "                        " + "\033[31m" + "▒" + "\033[0m" + "  " + "\033[31m" + "▒" + "\033[0m" + "   ",
		"             " + "\033[31m" + "▒" + "\033[0m" + "                                      ",
		"                                                    ",
		"                                                    ",
	}
	DisplayMenu(frame8)
	time.Sleep(250 * time.Millisecond)
	Clear()
	frame9 := []string{
		"█     █" + "\033[31m" + "░" + "\033[0m" + " ▄▄▄        ██████ ▄▄▄█████" + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "█████ " + "\033[31m" + "▓" + "\033[0m" + "█████▄  ",
		"" + "\033[31m" + "▓" + "\033[0m" + "█" + "\033[31m" + "░" + "\033[0m" + " █ " + "\033[31m" + "░" + "\033[0m" + "█" + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "████▄    " + "\033[31m" + "▒" + "\033[0m" + "██    " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "  ██" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "█     " + "\033[31m" + "▒" + "\033[0m" + "██  ██▌",
		"" + "\033[31m" + "▒" + "\033[0m" + "█" + "\033[31m" + "░" + "\033[0m" + " █ " + "\033[31m" + "░" + "\033[0m" + "█ " + "\033[31m" + "▒" + "\033[0m" + "██  ▀█▄  " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "██▄   " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "██" + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "███   " + "\033[31m" + "░" + "\033[0m" + "██   █▌",
		"" + "\033[31m" + "░" + "\033[0m" + "█" + "\033[31m" + "░" + "\033[0m" + " █ " + "\033[31m" + "░" + "\033[0m" + "█ " + "\033[31m" + "░" + "\033[0m" + "██▄▄▄▄██   " + "\033[31m" + "▒" + "\033[0m" + "   ██" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "██" + "\033[31m" + "▓" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "█    " + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "█    ▌",
		"" + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + "██" + "\033[31m" + "▒" + "\033[0m" + "██" + "\033[31m" + "▓" + "\033[0m" + "  " + "\033[31m" + "▓" + "\033[0m" + "█   " + "\033[31m" + "▓" + "\033[0m" + "██" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "██████" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "  " + "\033[31m" + "▒" + "\033[0m" + "██" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "████" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "█████▌ ",
		"" + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "   " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "   " + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "█" + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "  " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + "   " + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "  " + "\033[31m" + "▒" + "\033[0m" + " ",
		"" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "    " + "\033[31m" + "▒" + "\033[0m" + "   " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "   " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "    " + "\033[31m" + "░" + "\033[0m" + "       " + "\033[31m" + "░" + "\033[0m" + "  " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "  " + "\033[31m" + "▒" + "\033[0m" + "   ",
		"         " + "\033[31m" + "░" + "\033[0m" + "   " + "\033[31m" + "▒" + "\033[0m" + "   " + "\033[31m" + "░" + "\033[0m" + "                                  ",
		"                                                    ",
		"                                                    ",
	}
	DisplayMenu(frame9)
	time.Sleep(250 * time.Millisecond)
	Clear()
	frame10 := []string{
		"█     █" + "\033[31m" + "░" + "\033[0m" + " ▄▄▄        ██████ ▄▄▄█████" + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "█████ " + "\033[31m" + "▓" + "\033[0m" + "█████▄  ",
		"" + "\033[31m" + "▓" + "\033[0m" + "█" + "\033[31m" + "░" + "\033[0m" + " █ " + "\033[31m" + "░" + "\033[0m" + "█" + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "████▄    " + "\033[31m" + "▒" + "\033[0m" + "██    " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "  ██" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "█     " + "\033[31m" + "▒" + "\033[0m" + "██  ██▌",
		"" + "\033[31m" + "▒" + "\033[0m" + "█" + "\033[31m" + "░" + "\033[0m" + " █ " + "\033[31m" + "░" + "\033[0m" + "█ " + "\033[31m" + "▒" + "\033[0m" + "██  ▀█▄  " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "██▄   " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "██" + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "███   " + "\033[31m" + "░" + "\033[0m" + "██   █▌",
		"" + "\033[31m" + "░" + "\033[0m" + "█" + "\033[31m" + "░" + "\033[0m" + " █ " + "\033[31m" + "░" + "\033[0m" + "█ " + "\033[31m" + "░" + "\033[0m" + "██▄▄▄▄██   " + "\033[31m" + "▒" + "\033[0m" + "   ██" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "██" + "\033[31m" + "▓" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "█    " + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "█    ▌",
		"" + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + "██" + "\033[31m" + "▒" + "\033[0m" + "██" + "\033[31m" + "▓" + "\033[0m" + "  " + "\033[31m" + "▓" + "\033[0m" + "█   " + "\033[31m" + "▓" + "\033[0m" + "██" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "██████" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "  " + "\033[31m" + "▒" + "\033[0m" + "██" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "████" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "█████▌ ",
		"" + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "   " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "   " + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "█" + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "  " + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + "   " + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▓" + "\033[0m" + "  " + "\033[31m" + "▒" + "\033[0m" + " ",
		"" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "    " + "\033[31m" + "▒" + "\033[0m" + "   " + "\033[31m" + "▒" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "" + "\033[31m" + "▒" + "\033[0m" + "  " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "    " + "\033[31m" + "░" + "\033[0m" + "     " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "  " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "▒" + "\033[0m" + "  " + "\033[31m" + "▒" + "\033[0m" + "   ",
		"" + "\033[31m" + "░" + "\033[0m" + "   " + "\033[31m" + "░" + "\033[0m" + "    " + "\033[31m" + "░" + "\033[0m" + "   " + "\033[31m" + "▒" + "\033[0m" + "   " + "\033[31m" + "░" + "\033[0m" + "  " + "\033[31m" + "░" + "\033[0m" + "  " + "\033[31m" + "░" + "\033[0m" + "    " + "\033[31m" + "░" + "\033[0m" + "         " + "\033[31m" + "░" + "\033[0m" + "    " + "\033[31m" + "░" + "\033[0m" + " " + "\033[31m" + "░" + "\033[0m" + "  " + "\033[31m" + "░" + "\033[0m" + "   ",
		"" + "\033[31m" + "░" + "\033[0m" + "          " + "\033[31m" + "░" + "\033[0m" + "  " + "\033[31m" + "░" + "\033[0m" + "      " + "\033[31m" + "░" + "\033[0m" + "              " + "\033[31m" + "░" + "\033[0m" + "  " + "\033[31m" + "░" + "\033[0m" + "   " + "\033[31m" + "░" + "\033[0m" + "        ",
		"" + "\033[31m" + "░" + "\033[0m" + "                                                   ",
	}
	DisplayMenu(frame10)
	time.Sleep(250 * time.Millisecond)
	Write("\n\n\nClick on anything to continue...")
	anything := ""
	fmt.Scanln(&anything)
	p.Health_point = int(p.Max_health_point / 2)
	MainMenu(p)
}