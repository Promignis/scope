FROM golang:alpine as builder

COPY . $GOPATH/src/github.com/promignis/scope/
WORKDIR $GOPATH/src/github.com/promignis/scope/

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/scope

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/scope /go/bin/scope
ENV PORT 80
ENV APP_ENV PROD

ENTRYPOINT ["/go/bin/scope"]
