# syntax=docker/dockerfile:1

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.17-alpine as builder
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk add --no-cache git
WORKDIR $GOPATH/src/github.com/diegolopezcode/api-crud-complete-chi
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go install $GOPATH/src/github.com/diegolopezcode/api-crud-complete-chi/cmd

# FROM scratch

# COPY --from=builder /go/bin/cmd .
# EXPOSE 8080
# ENTRYPOINT ["./cmd"]

FROM alpine:3.14 as production
# Add certificates
RUN apk add --no-cache ca-certificates
# Copy built binary from builder
COPY --from=builder /go/bin/cmd .
# Expose port
EXPOSE 8080
ENTRYPOINT ["./cmd"]
