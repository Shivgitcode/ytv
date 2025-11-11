package main

import (
	"fmt"
	"os"
	"shivgitcode/youtubeview/cmd"
)


func main(){
	if len(os.Args)<2{
		fmt.Print("The list of available commands are\nstream\ndownload")
		return
	}

	command:=os.Args[1]
	args:=os.Args[2:]
	
	switch command{
	case "stream":
		cmd.StreamVideo(args)
	case "download":
		cmd.SaveVideo(args)
	}

}