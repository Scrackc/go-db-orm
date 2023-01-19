# build stage
FROM golang:1.19 as builder
WORKDIR /app
COPY ["go.mod", "go.sum", "./"]
RUN go mod download
COPY . .
RUN echo hola
RUN go build -o app -v .



# final stage
FROM alpine:3.17.1 as prod
WORKDIR /app
COPY --from=builder /app .
RUN apk add libc6-compat
ENTRYPOINT [ "./app" ]
