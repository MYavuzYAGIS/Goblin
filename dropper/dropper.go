package dropper

import (
	"fmt"

	"github.com/fatih/color"
)

func Dropper(url string) {
	fmt.Println(url)
	color.Red(url)
}
