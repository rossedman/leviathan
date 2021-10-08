# builder image
FROM registry.twilio.com/library/golang:1.16.2-1 as builder
WORKDIR /go/src/github.com/rossedman/leviathan
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
ENTRYPOINT ["leviathan"]