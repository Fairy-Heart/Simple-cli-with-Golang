package command

import (
	"fmt"
	"os"
)

func Exit() {
	fmt.Println(red_color("+ Exit Terminal"))
	os.Exit(1)
}
