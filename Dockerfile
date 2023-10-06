FROM golang:latest as builder

WORKDIR /app

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o api main.go

FROM scratch

COPY --from=builder /app /

CMD ["./api"]