package cmd

import (
	"flag"
	"fmt"
	"shivgitcode/youtubeview/internals"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)


var speeds=[]string{"1.25","1.5","1.75","2.0"}

func StreamVideo(values []string) {
	streamCmd := flag.NewFlagSet("stream", flag.ExitOnError)
    speedFlag := streamCmd.Bool("speed", false, "choose playback speed")

    streamCmd.Parse(values)
    rest := streamCmd.Args()

    if len(rest) == 0 {
        fmt.Println("stream requires a video URL")
        return
    }

    count:=0

    if strings.HasPrefix(rest[len(rest)-1],"--"){
        fmt.Println("flag should be before the URL")
        return
    }


    for _,arg:=range rest{
        if strings.HasPrefix(arg,"http"){
            count++
        }
        

    }
    
    if count>1{
        fmt.Println("stream command only accepts one video URL at a time")
        return
    }

    fmt.Println(*speedFlag)
    url := rest[0]
    speed := "1.0"

    if *speedFlag {
        selector := &survey.Select{
            Message: "Choose playback speed",
            Options: speeds,
        }
        survey.AskOne(selector, &speed)
    }

    internals.FetchVideoAndParse(url, speed)




}