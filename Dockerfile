FROM golang:1.11 AS builder

WORKDIR $GOPATH/src/bitbucket.org/joisandresky/qurban-v2

COPY . .

RUN go get -d -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/qurban_v2 .

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/qurban_v2 .

EXPOSE 8080

CMD ["./qurban_v2"]