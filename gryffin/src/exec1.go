package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	done := make(chan bool)
	cmd := exec.Command("./bin/render", "https://www.flickr.com/")
	// equivalent as running: render.js https://www.flickr.com/
	cmd.Stdout = os.Stdout
	cmd.Start()

	timeout := func() {
		<-time.After(time.Duration(15) * time.Second)
		cmd.Process.Kill()
		log.Println("Stop the renderer after 15 seconds for quick demo.")
		done <- true
	}
	go timeout()
	<-done
}
