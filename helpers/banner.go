package helpers

import (
	"fmt"

	"github.com/fatih/color"
)

const ASCII_BANNER = `
		 ______  _____  ______         _____ __   _
		|  ____ |     | |_____] |        |   | \  |
		|_____| |_____| |_____] |_____ __|__ |  \_|										   
`

const (
	SUB_TEXT string = "	A reverse shell and linux privesc automation tool. **MYavuzYAGIS**  **github.com/MYavuzYAGIS/Goblin"
)

func Print_Banner() {
	fmt.Println(ASCII_BANNER)
	fmt.Print("\n")
	color.Green(SUB_TEXT)
}
