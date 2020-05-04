FROM golang:1.14-stretch
ENV CGO_ENABLED=0
WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v

FROM alpine:3.11.6
COPY --from=0 /go/src/kpt-remove-resource /kpt-remove-resource
CMD ["/kpt-remove-resource"]
