package command

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func DownloadFile() {
	fmt.Print(yellow_color("Please input url to install file\n-> "))
	reader := bufio.NewReader(os.Stdin)
	url_file, err := reader.ReadString('\n')
	url_file = strings.TrimSpace(url_file)

	if err != nil {
		fmt.Println(red_color("The error has occurred in your command\nERROR -> ", err))
		return
	}

	res, err := http.Get(url_file)
	if err != nil {
		fmt.Println(red_color("The error has occurred in your command\nERROR -> ", err))
		return
	}
	defer res.Body.Close()

	fmt.Println(blue_color("Downloading files for you...\n"), yellow_color("-> Download from : "), url_file)

	file_name := path.Base(res.Request.URL.String())

	out, err := os.Create(file_name)
	if err != nil {
		fmt.Println(red_color("The error has occurred in your command\nERROR -> ", err))
		return
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		fmt.Println(red_color("The error has occurred in your command\nERROR -> ", err))
		return
	}
	fmt.Println(yellow_color("Download file successfuly with file name :\n-> ", file_name))
}
