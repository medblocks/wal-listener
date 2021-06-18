FROM golang:1.16 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./wal-listner ./cmd/wal-listener


FROM alpine  
WORKDIR /root/
COPY --from=build /src/wal-listner /wal-listner
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

CMD sh -c "/wait && /wal-listner"
