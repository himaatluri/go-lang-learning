package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

type testReport struct {
	Latency  string
	Download string
	Upload   string
	// Jitter   string
}

func main() {
	fmt.Println("Run speed test")
	op, err := exec.Command("/usr/bin/speedtest", "--accept-license", "--accept-gdpr").Output()

	if err != nil {
		log.Fatal(err)
	}
	//Latency
	lat, _ := regexp.Compile("Latency:\\s+(.*?)s")
	dwn, _ := regexp.Compile("Download:\\s+(.*?)s")
	upl, _ := regexp.Compile("Upload:\\s+(.*?)s")
	// jit, _ := regexp.Compile("Latency:.*?jitter:\\s+(.*?)ms")

	Latency := lat.FindString(string(op))
	Download := dwn.FindString(string(op))
	Upload := upl.FindString(string(op))
	// Jitter := jit.FindString(string(op))

	out := fmt.Sprintf("%s %s %s", Latency, Download, Upload)

	fmt.Println(out)

	rpt := testReport{Latency: Latency}
	rpt.Download = Download
	rpt.Upload = Upload
	// rpt.Jitter = Jitter

	fmt.Println(rpt)
}
