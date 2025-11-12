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

	"github.com/AlecAivazis/survey/v2"
	"github.com/Shivgitcode/ytv/internals"

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
	qualityFlag := saveCmd.Bool("quality",false,"this is to select download quality")

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
	quality:="360p"

	if *qualityFlag{
		options:=&survey.Select{
			Message: "Choose quality",
			Options: []string{
				"360p",
				"480p",
				"720p",
				"1080p",
			},

		}
		survey.AskOne(options,&quality)
	}

	cm2:=exec.Command("yt-dlp","--print","filename","-o",correctFilePath,videoUrl)

	filename,_:=cm2.Output()

	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("\x1b[34mDownloading...\x1b[0m"),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionThrottle(100*time.Millisecond),
		progressbar.OptionClearOnFinish(),
		
	)

	// ✅ Start yt-dlp
	cm := exec.Command("yt-dlp",
		"--newline",
		"--progress",
		"-f", internals.QualityToFormat(*qualityFlag,quality),
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
