FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/network-manager cmd/network-manager/main.go


FROM scratch AS runtime

COPY --from=builder /go/bin/network-manager /network-manager

ENTRYPOINT ["/network-manager"]
