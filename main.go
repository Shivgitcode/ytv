package main

import (
	"fmt"
	"os"
	"github.com/Shivgitcode/ytv/cmd"
)


func main(){
	if len(os.Args)<2 && os.Args[1]!="ytv"{
		fmt.Print("The list of available commands are\nstream\ndownload")
		return
	}
	if os.Args[1]=="ytv"{
		fmt.Println("Welcome to ytv")
		fmt.Println("Basic Commands ")
		fmt.Println("ytv stream <videourl> ")
		fmt.Println("ytv stream --speed <videourl>")
		fmt.Println("ytv download <videourl>")
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