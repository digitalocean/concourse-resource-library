package main

import (
	"encoding/json"
	"log"

	rlog "github.com/digitalocean/concourse-resource-library/log"
	resource "{{ .Module }}"
)

func main() {
	input := rlog.WriteStdin()
	defer rlog.Close()

	var request rshared.CheckRequest
	if err := json.Unmarshal(input, &request); err != nil {
		log.Fatalf("failed to unmarshal request: %s", err)
	}

	log.Println("check complete")
}
