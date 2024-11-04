FROM golang:1.21.3 as builder
ARG SERVICE_NAME
WORKDIR /src
COPY . .
RUN go build  -v -o /app ./cmd/${SERVICE_NAME}/...

FROM alpine:${ALPINE_VER}
RUN apt update && apt install -y ca-certificates && update-ca-certificates
COPY --from=builder /app /app
CMD ["/app"]
