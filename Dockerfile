FROM gcr.io/distroless/base

COPY wal-listener .

ENTRYPOINT ["./wal-listener"]