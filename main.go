package main

import (
	"path/filepath"
	"strings"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
	"github.com/sqweek/dialog"
)

const outFormat = "mp4"

func main() {

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
		panic(err.Error())
	}
}
