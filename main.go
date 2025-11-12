package main

import (
	"fmt"
	"os"

	"github.com/Shivgitcode/ytv/cmd"
	"github.com/common-nighthawk/go-figure"
)


func main() {
	if len(os.Args) < 2 {
		figure.NewFigure("YTV", "slant", true).Print()
		fmt.Println("Welcome to YTV - Your CLI YouTube Companion!")
		fmt.Println("\nUsage: ytv <command> [arguments]")
		fmt.Println("\nAvailable commands:")
		fmt.Println("  stream   : Stream a YouTube video.")
		fmt.Println("  download : Download a YouTube video.")
		fmt.Println("  playlist : Download a YouTube playlist.")
		return
	}

	command:=os.Args[1]
	args:=os.Args[2:]
	
	switch command{
	case "stream":
		cmd.StreamVideo(args)
	case "download":
		cmd.SaveVideo(args)
	case "playlist":
		cmd.SavePlaylist(args)
	default:
		fmt.Println("availaible commands\nstream\ndownload\nplaylist")
	}


}