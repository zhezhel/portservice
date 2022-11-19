FROM golang:alpine AS builder

COPY ./ /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /portservice -ldflags='-w -s -buildid= -extldflags "-static"' -tags timetzdata -trimpath ./cmd/portservice/main.go

FROM scratch
COPY --from=builder /portservice /bin/portservice
CMD ["/bin/portservice"]
