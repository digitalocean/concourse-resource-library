FROM golang:1.14 as builder
ADD . /go/src/{{ .Module }}
WORKDIR /go/src/{{ .Module }}
RUN make build

FROM alpine:3.11 as resource
RUN apk add --update --no-cache bash bash-completion openssh git
COPY --from=builder /go/src/{{ .Module }}/build /opt/resource
RUN ln -s /opt/resource/get /opt/resource/in && ln -s /opt/resource/put /opt/resource/out && chmod +x /opt/resource/*
CMD ["/bin/bash"]

FROM resource
LABEL MAINTAINER=digitalocean
