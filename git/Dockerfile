FROM alpine:3.10 as resource
LABEL MAINTAINER=digitalocean

RUN apk add --update --no-cache \
    git \
    openssh \
    && chmod +x /opt/resource/*
ADD scripts/install_git_crypt.sh install_git_crypt.sh
RUN ./install_git_crypt.sh && rm ./install_git_crypt.sh
