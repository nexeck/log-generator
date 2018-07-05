FROM golang:1.10-alpine

RUN mkdir -p /go/src/github.com/nexeck/log-generator
WORKDIR /go/src/github.com/nexeck/log-generator
COPY . .

RUN apk --no-cache add git
RUN go get -u -v github.com/golang/dep/cmd/dep
RUN dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app .

FROM alpine:latest
COPY --from=0 /app /app
ENTRYPOINT ["/app"]
