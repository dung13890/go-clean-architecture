FROM golang:1.20.3-alpine3.16 as base

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache gcc musl-dev make git bash openssh

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

# install the air binary so we get live code-reloading when we save files
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install github.com/dung13890/go-base-gen@latest
