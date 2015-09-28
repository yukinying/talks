package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	done := make(chan bool)
	cmd := exec.Command("./bin/using.arachni")
	// this is equivalent as running the above shell script.
	cmd.Stdout = os.Stdout
	cmd.Start()

	timeout := func() {
		<-time.After(time.Duration(15) * time.Second)
		cmd.Process.Kill()
		exec.Command("pkill", "-f", "fuzz-arachni")
		exec.Command("pkill", "-f", "arachni")
		log.Println("Stop the crawl after 15 second for quick demo")
		done <- true
	}
	go timeout()
	<-done
}
