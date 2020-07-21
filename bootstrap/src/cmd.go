package main

import (
	"log"
	{{- if ne .Cmd "Check" }}"os"{{ end -}}

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

	err = request.Source.Validate()
	if err != nil {
		log.Fatalf("invalid source config: %s", err)
	}

	{{ if ne .Cmd "Check" }}
	if len(os.Args) < 2 {
		log.Fatalf("missing arguments")
	}
	dir := os.Args[1]
	{{ end }}

	response, err := resource.{{ .Cmd }}(request{{ if ne .Cmd "Check" }}, dir{{ end }})
	if err != nil {
		log.Fatalf("failed to perform check: %s", err)
	}

	{{ if eq .Cmd "Get" }}
	// write metadata to output dir
	err = response.Metadata.ToFiles(filepath.Join(dir, "resource"))
	if err != nil {
		log.Fatalf("failed to write metadata.json: %s", err)
	}
	{{ end }}

	err = response.Write()
	if err != nil {
		log.Fatalf("failed to write response to stdout: %s", err)
	}

	log.Println("{{ lower .Cmd }} complete")
}
