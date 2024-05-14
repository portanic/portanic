FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /portanic ./cmd/portanic

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /portanic .

RUN mkdir /usr/share/portanic/

COPY static /root/static

CMD ["./portanic"]
