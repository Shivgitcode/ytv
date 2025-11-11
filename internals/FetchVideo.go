package internals

import (
	"fmt"

	"os/exec"
)


func checkIfPresent(bin string) bool{
	_,err:=exec.LookPath(bin)

	return err==nil
}

func FetchVideoAndParse(url string,speed string) {

	if !checkIfPresent("yt-dlp"){
		fmt.Println("yt-dlp is not present first install that")
		return
	}
	if !checkIfPresent("mpv"){
		fmt.Println("mpv is not present first install that")
		return
	}
	
	videoURL := url



	cmd:=exec.Command("yt-dlp" ,"-f", "best", "-g",videoURL)
	out,err:=cmd.Output()
	if err!=nil {
		panic(err)
	}

	formatStr:=fmt.Sprintf("--speed=%s",speed)


	playCommand:=exec.Command("mpv","--focus-on=open",formatStr,string(out))

	if err:=playCommand.Run();err!=nil{
		fmt.Println("mpv error")
		panic(err)
	}



	




}
