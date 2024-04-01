FROM golang:1.15
WORKDIR /go/src/victorapdias
COPY . .
RUN GOSS=linux go build -ldflags="-s -w"
CMD ["./victorapdias"]