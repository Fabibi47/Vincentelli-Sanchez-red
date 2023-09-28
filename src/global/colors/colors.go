package colors

import "runtime"

var Reset = "\033[0m"
var Red = "\033[31m"
var LightRed = "\033[91m"
var Green = "\033[32m"
var LightGreen = "\033[92m"
var Yellow = "\033[33m"
var LightYellow = "\033[93m"
var Blue = "\033[34m"
var LightBlue = "\033[94m"
var Purple = "\033[35m"
var LightMagenta = "\033[95m"
var Cyan = "\033[36m"
var LightCyan = "\033[96m"
var LightGray = "\033[37m"
var DarkGray = "\033[90m"
var White = "\033[97m"

func Init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		LightRed = ""
		Green = ""
		LightGreen = ""
		Yellow = ""
		LightYellow = ""
		Blue = ""
		LightBlue = ""
		Purple = ""
		LightMagenta = ""
		Cyan = ""
		LightCyan = ""
		DarkGray = ""
		LightGray = ""
		White = ""
	}
}
