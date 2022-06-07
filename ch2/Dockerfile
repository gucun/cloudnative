FROM golang:1.18 as builder
ENV GOPROXY=https://goproxy.cn
ENV CGO_ENABLED=0
ENV GO111MODULE=on

WORKDIR /usr/src/app
COPY . .

RUN go build  -o server

# FROM scratch as runner
# COPY --from=builder /usr/src/app/server /opt/app/
# CMD ["/opt/app/server"]


FROM alpine:latest

WORKDIR /usr/local/bin/

COPY --from=builder /usr/src/app/server .
CMD ["./server"]
# ENTRYPOINT [ "/server" ]
