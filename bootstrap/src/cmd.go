package main

import (
	"log"

	rlog "github.com/digitalocean/concourse-resource-library/log"
	resource "{{ .Module }}"
)

func main() {
	input := rlog.WriteStdin()
	defer rlog.Close()

	var request resource.{{ .Cmd }}Request
	err := request.Read(input)
	if err != nil {
		log.Fatalf("failed to read request input: %s", err)
	}

	response, err := resource.{{ .Cmd }}(request)
	if err != nil {
		log.Fatalf("failed to perform check: %s", err)
	}

	err = response.Write()
	if err != nil {
		log.Fatalf("failed to write response to stdout: %s", err)
	}

	log.Println("{{ .Cmd }} complete")
}
