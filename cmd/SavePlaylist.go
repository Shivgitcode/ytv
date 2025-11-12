package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Shivgitcode/ytv/internals"
	"github.com/schollz/progressbar/v3"
)



func SavePlaylist(values []string){
	savePlaylistCmd:=flag.NewFlagSet("playlist",flag.ExitOnError)

	err:=savePlaylistCmd.Parse(values)

	internals.Check(err)

	rest:=savePlaylistCmd.Args()
	playlistUrl:=rest[0]
	homePath,_:=os.UserHomeDir()
	completePath:=filepath.Join(homePath,"Downloads/%(playlist_title)s/%(playlist_index)03d - %(title)s.%(ext)s")
	if len(rest) < 1 {
		fmt.Println("provide a url")
		return
	}

	count:=0
	for _,val:=range rest{
		if strings.HasPrefix("https",val){
			count++
		}
	}
	if count>1{
		fmt.Println("Only one Playlist can be downloaded at a time")
		return
	}

	bar2 := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("\x1b[34mPreparing to download\x1b[0m"),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionThrottle(100*time.Millisecond),
		progressbar.OptionClearOnFinish(),
		
	)

	go func(){
		for {
			bar2.Add(1)
			time.Sleep(120*time.Millisecond)
		}

	}()

	cmd2:=exec.Command("yt-dlp","--print","filename","-o",completePath,playlistUrl)

	filename,_:=cmd2.Output()

	

	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("\x1b[34mDownloading...\x1b[0m"),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionThrottle(100*time.Millisecond),
		progressbar.OptionClearOnFinish(),
		
	)
	

	cmd:=exec.Command("yt-dlp","-f","bv*[ext=mp4]+ba[ext=m4a]/b[ext=mp4]","--yes-playlist","--concurrent-fragments","4","--merge-output-format","mp4","-o",completePath,playlistUrl)


	
	stdout,_:=cmd.StdoutPipe()
	cmd.Start()

	startRealDownload:=false
	go func(){
		scanner:=bufio.NewScanner(stdout)
		for scanner.Scan(){
			if !startRealDownload{
				startRealDownload=true
				bar2.Finish()

				fmt.Println()
			}
			bar.Add(1)




		}
	}()

	cmd.Wait()
	bar.Finish()

	fmt.Println("Playlist Downloaded successfully at",string(filename))


	







	

}