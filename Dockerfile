FROM golang:alpine AS builder

RUN addgroup -S portservice && adduser -S -u 10000 -g portservice portservice

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o /portservice -ldflags='-w -s -buildid= -extldflags "-static"' \
    -tags timetzdata -trimpath ./cmd/portservice/

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /portservice /bin/portservice

USER portservice:portservice

CMD ["/bin/portservice"]
