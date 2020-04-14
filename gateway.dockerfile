FROM golang:1.14.1-alpine3.11 AS build
ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
WORKDIR /go/release
ADD . .
RUN go mod download  \
    && GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o main gateway/main.go

FROM alpine:3.11.5 as prod

COPY --from=build /go/release/main /

CMD ["/main"]