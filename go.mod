module github.com/digitalocean/concourse-resource-library

go 1.14

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/Microsoft/go-winio v0.4.15-0.20190919025122-fc70bd9a86b5 // indirect
	github.com/VividCortex/ewma v1.1.1 // indirect
	github.com/concourse/go-archive v1.0.1
	github.com/containerd/containerd v1.3.6 // indirect
	github.com/docker/docker v1.13.1
	github.com/fatih/color v1.9.0
	github.com/go-git/go-git/v5 v5.0.0
	github.com/google/go-containerregistry v0.1.1
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/huandu/xstrings v1.3.1 // indirect
	github.com/jfrog/jfrog-client-go v0.12.0
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/moby/term v0.0.0-20200611042045-63b9a826fb74 // indirect
	github.com/nelsam/hel v0.0.0-20200611165952-2d829bae0c66 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/poy/onpar v0.0.0-20200406201722-06f95a1c68e8
	github.com/shurcooL/githubv4 v0.0.0-20191127044304-8f68eb5628d0
	github.com/shurcooL/graphql v0.0.0-20181231061246-d48a9a75455f // indirect
	github.com/spf13/cobra v1.0.0
	github.com/vbauerster/mpb v3.4.0+incompatible
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
)

replace github.com/docker/docker v1.13.1 => github.com/docker/engine v0.0.0-20200720230453-22153d111ead
