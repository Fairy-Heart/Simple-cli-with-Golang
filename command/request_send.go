package command

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

)

func Request_Send() {
	fmt.Print(yellow_color("Please input url to make HTTP request\n-> "))
	reader := bufio.NewReader(os.Stdin)
	cmds, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(red_color("The error has occurred in your command\nERROR -> ", err))
		return
	}

	cmds = strings.TrimSpace(cmds)
	request_url := cmds

	res, err := http.Get(request_url)

	if err != nil {
		fmt.Println(red_color("The error has occurred in your command\nERROR -> ", err))
		return
	}
	Start_Time := time.Now()
	time.Sleep(25 * time.Second)
	Current_Time := time.Now()
	Time_Out := Current_Time.Sub(Start_Time)

	if Time_Out >= 25 {
		fmt.Println(red_color("TIME OUT !!!\n-> The website you want to send a request to did not respond within 25 seconds."))
		return
	}

	defer res.Body.Close()
	fmt.Println(yellow_color("Send request to ", request_url, " successfuly..\nStatus code : ", res.StatusCode))
}
