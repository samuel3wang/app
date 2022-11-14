FROM golang:alpine AS builder
ADD . /src
RUN cd /src \
    && go build -o app

FROM alpine
WORKDIR /app
COPY --from=builder /src/app .
CMD ["./app"]