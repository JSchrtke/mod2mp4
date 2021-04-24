package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
	"github.com/sqweek/dialog"
)

const outFormat = "mp4"

func main() {
	ffmpegPath, err := getFfmpegPath()
	if err != nil {
		dialog.Message(fmt.Sprintf("Missing ffmpeg installation: %s", err.Error())).
			Title("Error").
			Info()
	}

	filename, err := dialog.File().Filter("MOD video file", "MOD").Load()
	if err == dialog.Cancelled {
		return
	} else if err != nil {
		panic(err)
	}

	outFile, err := dialog.File().Filter("mp4 video file", "mp4").Save()
	if err == dialog.Cancelled {
		return
	} else if err != nil {
		panic(err)
	}
	if !strings.HasSuffix(filepath.Ext(outFile), "mp4") {
		outFile = strings.Join([]string{outFile, "mp4"}, ".")
	}

	err = fluentffmpeg.NewCommand(ffmpegPath).
		InputPath(filename).
		OutputFormat(outFormat).
		OutputPath(outFile).
		Overwrite(true).
		Run()
	if err != nil {
		dialog.Message(err.Error()).
			Title(fmt.Sprintf("error converting %s", filename)).
			Info()
	}
}

func getFfmpegPath() (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	var ffmpegPath string
	if runtime.GOOS == "windows" {
		ffmpegPath = filepath.Clean(filepath.Join(workingDir, "deps/ffmpeg-win64/bin/ffmpeg.exe"))
	} else if runtime.GOOS == "linux" {
		ffmpegPath = "deps/ffmpeg-linux64/ffmpeg"
	} else {
		ffmpegPath = "ffmpeg"
	}
	_, err = exec.LookPath(ffmpegPath)
	if err != nil {
		return "", err
	}

	return ffmpegPath, nil
}
