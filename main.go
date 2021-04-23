package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
	"github.com/sqweek/dialog"
)

const outFormat = "mp4"

func main() {
	ffmpegPath := "ffmpeg"
	_, err := exec.LookPath(ffmpegPath)
	if err != nil {
		if runtime.GOOS == "windows" {
			ffmpegPath = "./deps/ffmpeg-win64/bin/ffmpeg.exe"
		} else {
			ffmpegPath = "."
		}
		// check if there are bundled versions
		_, err := exec.LookPath(ffmpegPath)
		if err != nil {
			dialog.Message("Missing ffmpeg installation").Title("Error").Info()
			return
		}
	}

	filename, err := dialog.File().Filter("MOD video file", "MOD").Load()
	if err != nil {
		if err == dialog.Cancelled {
			return
		} else {
			panic(err.Error())
		}
	}

	outFile := strings.TrimSuffix(filename, filepath.Ext(filename))
	outFile = strings.Join([]string{outFile, outFormat}, ".")

	err = fluentffmpeg.NewCommand("").
		InputPath(filename).
		OutputFormat(outFormat).
		OutputPath(outFile).
		Overwrite(true).
		Run()
	if err != nil {
		dialog.Message(err.Error()).Title(fmt.Sprintf("error converting %s", filename)).Info()
	}
}
