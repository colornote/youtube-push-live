package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

const (

	// ffmpegDarwin MacOS 下 ffmpeg 的位置
	ffmpegDarwin = "./ffmpeg/ffmpeg"
	// Linux 同理
	ffmpegLinux = "./ffmpeg/ffmpeg"
	// Windows 同理
	ffmpegWindows = "./ffmpeg/bin/ffmpeg.exe"
)

func main() {

	// read key from file
	url, err := ioutil.ReadFile("url.txt")

	if err != nil || len(url) == 0 {
		fmt.Println("key is required", err)
		return
	}

	// read all files in videos folder
	files, err := ioutil.ReadDir("./videos")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("开始推流...")

LOOP:
	for _, file := range files {
		push(string(url), "./videos/"+file.Name())
	}
	goto LOOP

}

func push(url, filename string) {

	// ffmpeg 在当前系统中的位置
	var ffmpeg string
	switch runtime.GOOS {
	case "darwin":
		ffmpeg = ffmpegDarwin
	case "windows":
		ffmpeg = ffmpegWindows
	case "linux":
		ffmpeg = ffmpegLinux
	}

	// cat *.mp4  | ffmpeg -re -i pipe:0 -c:v copy -c:a aac -f flv rtmp://a.rtmp.youtube.com/live2/xvkt-e9sp-ahh2-17za

	// 拼凑参数集合
	cmdArguments := []string{"-re", "-i", filename, "-c:v", "copy", "-c:a", "aac", "-f", "flv", strings.TrimSpace(string(url))}

	// 调用当前系统 cmd 执行 ffmpeg
	cmd := exec.Command(ffmpeg, cmdArguments...)

	// fmt.Println(cmd)
	// return

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
