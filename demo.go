//ffmpeg -i /go/src/go-ffmpeg/raw/ellen.mp4 -y -vf subtitles=/go/src/go-ffmpeg/raw/ellen.srt /go/src/go-ffmpeg/video/ellen.mp4
package main

import (
	"bufio"
	"errors"
	"log"
	//"os"
	"os/exec"
)

func runFfmpeg(str ...string) (errMess error) {
	cmd := exec.Command("ffmpeg", str...)
	_, err := cmd.StdoutPipe()
	checkError(err)
	stderr, err := cmd.StderrPipe()
	checkError(err)
	errBuf := bufio.NewReader(stderr)
	go func() {
		for {
			errLine, errErr := errBuf.ReadBytes('\n')
			if errErr != nil {
				break
			}
			errMess = errors.New(string(errLine))
		}
	}()
	cmd.Start()
	defer cmd.Wait()
	return errMess
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func main() {
	arg := []string{
		"-i",
		"/go/src/go-ffmpeg/raw/ellen.mp4",
		"-y",
		"-vf",
		"subtitles=/go/src/go-ffmpeg/raw/ellen.srt",
		"/go/src/go-ffmpeg/video/ellen.mp4",
	}
	if errRun := runFfmpeg(arg...); errRun != nil {
		log.Fatalf("Error: %s", errRun)
	}
}