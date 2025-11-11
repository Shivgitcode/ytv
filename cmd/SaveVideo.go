package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/shivgitcode/youtubeview/internals"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

type Progress struct {
	Downloaded int64 `json:"downloaded_bytes"`
	Total      int64 `json:"total_bytes"`
}

type YTDLPProgress struct {
	Progress Progress `json:"progress"`
}

func SaveVideo(values []string) {
	saveCmd := flag.NewFlagSet("download", flag.ExitOnError)
	qualityFlag := saveCmd.String("quality", "720p", "To tell the cli in which quality you want to stream the video")

	err := saveCmd.Parse(values)
	internals.Check(err)
	rest := saveCmd.Args()

	if len(rest) < 1 {
		fmt.Println("provide a url")
		return
	}

	videoUrl := rest[0]
	home, _ := os.UserHomeDir()

	correctFilePath := filepath.Join(home, "Downloads/%(title)s.%(ext)s")

	count := 0
	for _, val := range rest {
		if strings.HasPrefix(val, "http") {
			count++
		}
	}
	if count > 1 {
		fmt.Println("Only one Video can be downloaded at a time")
		return
	}

	if strings.HasPrefix(rest[len(rest)-1], "-") {
		fmt.Println("Flag should come before url")
		return
	}

	cm2:=exec.Command("yt-dlp","--print","filename","-o",correctFilePath,videoUrl)

	filename,_:=cm2.Output()

	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("Downloading..."),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionThrottle(100*time.Millisecond),
		progressbar.OptionClearOnFinish(),
	)

	// ✅ Start yt-dlp
	cm := exec.Command("yt-dlp",
		"--newline",
		"--progress",
		"-f", "best",
		"-o", correctFilePath,
		videoUrl,
	)


	stdout, _ := cm.StdoutPipe()
	cm.Start()


	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			bar.Add(1) 
		}
	}()

	cm.Wait()
	bar.Finish()



	fmt.Println("\nDownloaded on path",strings.TrimSpace(string(filename)))
	fmt.Println(*qualityFlag)
}
