package main

import (
	"Golang/command"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
)

var (
	reader *bufio.Reader
	cmds   string
	err    error
)

var (
	red_color   = color.Style{color.FgRed}.Render
	green_color = color.Style{color.FgGreen}.Render
	// yellow_color = color.Style{color.FgYellow}.Render
	// blue_color   = color.Style{color.FgBlue}.Render
)

func cmd() {
	color.Set(color.FgCyan)
	for {
		fmt.Print(green_color("$ "))
		reader = bufio.NewReader(os.Stdin)
		cmds, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(red_color(" The error has occurred in your command\nERROR -> ", err))
		}
		cmds = strings.TrimSpace(cmds)

		switch cmds {
		case "rqs", "request send", "make request", "mkreq":
			command.Request_Send()
		case "download file", "dfl", "df":
			command.DownloadFile()
		case "test":
			command.Test()
		case "exit":
			command.Exit()
		default:
			fmt.Println(red_color("+ The error has occurred in your command"))
			fmt.Println(red_color("+ ERROR -> COMMAND NOT FOUND"))
			fmt.Println(red_color("^^ ", cmds))
		}
	}
}

func main() {
	cmd()
}
