##########################################################
# Dockerfile to build Omigad container images
# Based on scratch
##########################################################
#!/bin/bash
# FROM golang:1.10.1-alpine3.7
FROM scratch

LABEL maintainer="liuzh <liuzh@llvision.com>"

# ENV GOPATH /go
# ENV PATH $PATH:$GOPATH/bin
# ENV CGO_ENABLE 0

# WORKDIR /go/src/github.com/wst-libs

ADD ./omigad /wst/omigad
ADD ./conf /wst/conf

WORKDIR /wst

EXPOSE 18012

CMD ["/wst/omigad"]