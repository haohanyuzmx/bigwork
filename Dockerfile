#FROM golang:latest
#
#WORKDIR $GOPATH/scr/hello/bigwork
#COPY . $GOPATH/src/scr/hello/bigwork
#RUN go build cmd/main.go
#
#EXPOSE 8000
#ENTRYPOINT ["./main"]