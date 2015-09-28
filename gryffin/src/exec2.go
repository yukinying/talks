package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	done := make(chan bool, 10)
	cmd := exec.Command("./bin/gryffin.demo", "http://52.89.152.13:8080/xss/reflect/full1?in=x")
	// this is equivalent as running the above shell script.
	cmd.Stdout = os.Stdout
	cmd.Start()
	pid := cmd.Process.Pid

	timeout := func() {
		<-time.After(time.Duration(5) * time.Second)
		cmd.Process.Kill()
		exec.Command("kill", fmt.Sprintf("%d", pid))
		exec.Command("pkill", "-f", "exe/main crawl")
		exec.Command("pkill", "-f", "phantomjs")
		log.Println("Stop the crawl after 5 seconds for quick demo.")
		done <- true
	}
	go timeout()
	<-done
	cmd.Wait()

}
