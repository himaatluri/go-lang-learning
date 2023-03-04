package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Run speed test")
	op, err := exec.Command("/usr/bin/speedtest", "--accept-license", "--accept-gdpr").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(op))
}
