package internals

import (
	"fmt"
	"strconv"
	"strings"
)


func Check(e error){
	if e!=nil{
		panic(e)
	}
}

func QualityToFormat(q bool,option string) string{
	if !q{
		return "best"
	}
	option=strings.ToLower(strings.TrimSpace(option))
	option=strings.TrimSuffix(option,"p")

	height,err:=strconv.Atoi(option)
	if err!=nil{
		return "best"
	}

	return fmt.Sprintf("bestvideo[height=%d][ext=mp4][vcodec*=avc1]+bestaudio[ext=m4a][acodec*=mp4a]/best[height=%d][ext=mp4][vcodec*=avc1]", height, height)
}